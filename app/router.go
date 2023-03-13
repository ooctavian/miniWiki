package app

import (
	"net/http"

	"miniWiki/config"
	accController "miniWiki/domain/account/controller"
	accQuery "miniWiki/domain/account/query"
	accService "miniWiki/domain/account/service"
	auController "miniWiki/domain/auth/controller"
	auQuery "miniWiki/domain/auth/query"
	auService "miniWiki/domain/auth/service"
	cController "miniWiki/domain/category/controller"
	cQuery "miniWiki/domain/category/query"
	cService "miniWiki/domain/category/service"
	iService "miniWiki/domain/image/service"
	pController "miniWiki/domain/profile/controller"
	pQuery "miniWiki/domain/profile/query"
	pService "miniWiki/domain/profile/service"
	rController "miniWiki/domain/resource/controller"
	rQuery "miniWiki/domain/resource/query"
	rService "miniWiki/domain/resource/service"
	"miniWiki/middleware"
	"miniWiki/security"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitRouter(conn *pgxpool.Pool, cfg config.Config) http.Handler {
	categoryQuerier := cQuery.NewQuerier(conn)
	imageService := iService.NewImage(cfg.Database.ImageDir)
	resourceService := rService.NewResource(rQuery.NewQuerier(conn), categoryQuerier, imageService)
	categoryService := cService.NewCategory(categoryQuerier)
	argon2id := security.NewArgon2id(
		cfg.Argon2id.Memory,
		cfg.Argon2id.Iterations,
		cfg.Argon2id.Parallelism,
		cfg.Argon2id.SaltLength,
		cfg.Argon2id.KeyLength,
		security.GenerateRandomBytes,
	)
	accountService := accService.NewAccount(accQuery.NewQuerier(conn), argon2id)
	authQuerier := auQuery.NewQuerier(conn)
	authService := auService.NewAuth(
		authQuerier,
		accQuery.NewQuerier(conn),
		argon2id,
	)
	profileService := pService.NewProfile(pQuery.NewQuerier(conn), imageService)

	sessionMiddleware := middleware.SessionMiddleware(authQuerier)

	r := chi.NewRouter()
	r.Group(func(gr chi.Router) {
		gr.Route("/resources", func(rr chi.Router) {
			rr.Use(sessionMiddleware)
			rController.MakeResourceRouter(rr, resourceService)
		})
		gr.Route("/profile", func(rr chi.Router) {
			rr.Use(sessionMiddleware)
			pController.MakeProfileRouter(rr, profileService)
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
