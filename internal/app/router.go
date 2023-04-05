package app

import (
	"net/http"
	"strings"

	auController "miniWiki/internal/auth/controller"
	auRepository "miniWiki/internal/auth/repository"
	auService "miniWiki/internal/auth/service"
	"miniWiki/internal/config"
	accController "miniWiki/internal/domain/account/controller"
	aRepository "miniWiki/internal/domain/account/repository"
	accService "miniWiki/internal/domain/account/service"
	cController "miniWiki/internal/domain/category/controller"
	cRepository "miniWiki/internal/domain/category/repository"
	cService "miniWiki/internal/domain/category/service"
	iService "miniWiki/internal/domain/image/service"
	rController "miniWiki/internal/domain/resource/controller"
	rRepository "miniWiki/internal/domain/resource/repository"
	rService "miniWiki/internal/domain/resource/service"
	"miniWiki/internal/domain/swagger"
	"miniWiki/internal/middleware"
	"miniWiki/pkg/security"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, cfg config.Config) http.Handler {
	imageService := iService.NewImage(cfg.Database.ImageDir)
	resourceRepository := rRepository.NewResourceRepository(db)
	categoryService := cService.NewCategory(cRepository.NewCategoryRepository(db), resourceRepository)
	resourceService := rService.NewResource(resourceRepository, categoryService, imageService)
	argon2id := security.NewArgon2id(
		cfg.Argon2id.Memory,
		cfg.Argon2id.Iterations,
		cfg.Argon2id.Parallelism,
		cfg.Argon2id.SaltLength,
		cfg.Argon2id.KeyLength,
		security.GenerateRandomBytes,
	)
	accountRepository := aRepository.NewAccountRepository(db)
	authRepository := auRepository.NewAuthRepository(db)
	accountService := accService.NewAccount(accountRepository, resourceRepository, argon2id, imageService)
	authService := auService.NewAuth(
		accountRepository,
		authRepository,
		argon2id,
	)

	sessionMiddleware := middleware.SessionMiddleware(authService)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(cfg.Server.AllowOrigin, ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Cookie"},
		ExposedHeaders:   []string{"Link", "Location", "Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/swagger/*", swagger.Handler())
	fs := http.FileServer(http.Dir(cfg.Database.ImageDir))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Group(func(gr chi.Router) {
		gr.Route("/resources", func(rr chi.Router) {
			rr.Use(sessionMiddleware)
			rController.MakeResourceRouter(rr, resourceService)
		})
		gr.Route("/categories", func(cr chi.Router) {
			cr.Use(sessionMiddleware)
			cController.MakeCategoryRouter(cr, categoryService)
		})
		gr.Route("/account", func(ar chi.Router) {
			accController.MakeAccountRouter(ar, accountService)
			ar.Group(func(apr chi.Router) {
				apr.Use(sessionMiddleware)
				accController.MakePrivateAccountRouter(apr, accountService)
			})

		})
		gr.Route("/", func(ar chi.Router) {
			auController.MakeAuthRouter(ar, authService)
			ar.Group(func(sr chi.Router) {
				sr.Use(sessionMiddleware)
				auController.MakeProtectedAuthRouter(sr, authService)
			})
		})
	})
	return r
}
