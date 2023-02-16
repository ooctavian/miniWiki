package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	rController "miniWiki/domain/resource/controller"
	rQuery "miniWiki/domain/resource/query"
	rService "miniWiki/domain/resource/service"

	iController "miniWiki/domain/image/controller"
	iService "miniWiki/domain/image/service"

	cController "miniWiki/domain/category/controller"
	cQuery "miniWiki/domain/category/query"
	cService "miniWiki/domain/category/service"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func createDir(dirName string) {
	if _, err := os.Stat(dirName); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			utils.Logger.Fatalf("Creating dir error: %v", err)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Fatalf("Error loading .env file")
	}
	createDir(os.Getenv("IMAGES_DIR") + "/resources")
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	port := os.Getenv("PORT")
	resourceService := rService.NewResource(rQuery.NewQuerier(pool))
	imageService := iService.NewImage()
	categoryService := cService.NewCategory(cQuery.NewQuerier(pool))

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

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		utils.Logger.Fatalf("Error while starting server %v", err)
	}
}
