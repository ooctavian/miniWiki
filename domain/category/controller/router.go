package controller

import (
	"context"

	"miniWiki/domain/category/model"

	"github.com/go-chi/chi/v5"
)

type categoryService interface {
	CreateCategory(ctx context.Context, request model.CreateCategoryRequest) (*int, error)
	DeleteCategory(ctx context.Context, request model.DeleteCategoryRequest) error
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategory(ctx context.Context, request model.GetCategoryRequest) (*model.Category, error)
}

func MakeCategoryRouter(r chi.Router, service categoryService) {
	r.Post("/", createCategoryHandler(service))
	r.Delete("/{id}", deleteResourceHandler(service))
	r.Get("/{id}", getResourceHandler(service))
	r.Get("/", getResourcesHandler(service))
}
