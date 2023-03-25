package app

import (
	"net/http"

	"miniWiki/pkg/config"
	accController "miniWiki/pkg/domain/account/controller"
	accQuery "miniWiki/pkg/domain/account/query"
	accService "miniWiki/pkg/domain/account/service"
	auController "miniWiki/pkg/domain/auth/controller"
	auQuery "miniWiki/pkg/domain/auth/query"
	auService "miniWiki/pkg/domain/auth/service"
	cController "miniWiki/pkg/domain/category/controller"
	cService "miniWiki/pkg/domain/category/service"
	iService "miniWiki/pkg/domain/image/service"
	rController "miniWiki/pkg/domain/resource/controller"
	rQuery "miniWiki/pkg/domain/resource/query"
	rService "miniWiki/pkg/domain/resource/service"
	"miniWiki/pkg/domain/swagger"
	"miniWiki/pkg/middleware"
	security2 "miniWiki/pkg/security"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

func InitRouter(conn *pgxpool.Pool, db *gorm.DB, cfg config.Config) http.Handler {
	resourceQuerier := rQuery.NewQuerier(conn)
	imageService := iService.NewImage(cfg.Database.ImageDir)
	resourceService := rService.NewResource(resourceQuerier, imageService)
	categoryService := cService.NewCategory(db)
	argon2id := security2.NewArgon2id(
		cfg.Argon2id.Memory,
		cfg.Argon2id.Iterations,
		cfg.Argon2id.Parallelism,
		cfg.Argon2id.SaltLength,
		cfg.Argon2id.KeyLength,
		security2.GenerateRandomBytes,
	)
	accountService := accService.NewAccount(db, argon2id, imageService)
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
