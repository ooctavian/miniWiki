package app

import (
	"net/http"

	auController "miniWiki/internal/auth/controller"
	auQuery "miniWiki/internal/auth/query"
	auService "miniWiki/internal/auth/service"
	accController "miniWiki/internal/domain/account/controller"
	accQuery "miniWiki/internal/domain/account/query"
	aRepository "miniWiki/internal/domain/account/repository"
	accService "miniWiki/internal/domain/account/service"
	cController "miniWiki/internal/domain/category/controller"
	cRepository "miniWiki/internal/domain/category/repository"
	cService "miniWiki/internal/domain/category/service"
	iService "miniWiki/internal/domain/image/service"
	rController "miniWiki/internal/domain/resource/controller"
	rQuery "miniWiki/internal/domain/resource/query"
	rRepository "miniWiki/internal/domain/resource/repository"
	rService "miniWiki/internal/domain/resource/service"
	"miniWiki/internal/domain/swagger"
	"miniWiki/internal/middleware"
	"miniWiki/pkg/config"
	"miniWiki/pkg/security"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

func InitRouter(conn *pgxpool.Pool, db *gorm.DB, cfg config.Config) http.Handler {
	resourceQuerier := rQuery.NewQuerier(conn)
	imageService := iService.NewImage(cfg.Database.ImageDir)
	resourceService := rService.NewResource(resourceQuerier, imageService)
	resourceRepository := rRepository.NewResourceRepository(db)
	categoryService := cService.NewCategory(cRepository.NewCategoryRepository(db), resourceRepository)
	argon2id := security.NewArgon2id(
		cfg.Argon2id.Memory,
		cfg.Argon2id.Iterations,
		cfg.Argon2id.Parallelism,
		cfg.Argon2id.SaltLength,
		cfg.Argon2id.KeyLength,
		security.GenerateRandomBytes,
	)
	accountService := accService.NewAccount(aRepository.NewAccountRepository(db), resourceRepository, argon2id, imageService)
	authQuerier := auQuery.NewQuerier(conn)
	authService := auService.NewAuth(
		authQuerier,
		accQuery.NewQuerier(conn),
		argon2id,
	)

	sessionMiddleware := middleware.SessionMiddleware(authService)

	r := chi.NewRouter()
	r.Get("/swagger/*", swagger.Handler())
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
