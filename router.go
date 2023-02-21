package miniWiki

import (
	"context"
	"net/http"

	cController "miniWiki/domain/category/controller"
	cQuery "miniWiki/domain/category/query"
	cService "miniWiki/domain/category/service"
	iController "miniWiki/domain/image/controller"
	iService "miniWiki/domain/image/service"
	rController "miniWiki/domain/resource/controller"
	rQuery "miniWiki/domain/resource/query"
	rService "miniWiki/domain/resource/service"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Conn interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

func InitRouter(conn Conn) http.Handler {
	resourceService := rService.NewResource(rQuery.NewQuerier(conn))
	imageService := iService.NewImage()
	categoryService := cService.NewCategory(cQuery.NewQuerier(conn))

	r := chi.NewRouter()
	r.Group(func(gr chi.Router) {
		gr.Route(
			"/resources",
			func(rr chi.Router) {
				rController.MakeResourceRouter(rr, resourceService)
				iController.MakeResourceImageHandler(rr, imageService)
			})
		gr.Route(
			"/categories",
			func(cr chi.Router) {
				cController.MakeCategoryRouter(cr, categoryService)
			})
	})
	return r
}
