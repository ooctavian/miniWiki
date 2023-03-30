package controller

import (
	"context"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request model.CreateCategoryRequest) (*int, error)
	DeleteCategory(ctx context.Context, request model.DeleteCategoryRequest) error
	GetCategories(ctx context.Context, pagination utils.Pagination) (*utils.Pagination, error)
	GetCategory(ctx context.Context, request model.GetCategoryRequest) (*model.Category, error)
}

func MakeCategoryRouter(r chi.Router, service CategoryService) {
	r.Post("/", createCategoryHandler(service))
	r.Delete("/{id}", deleteResourceHandler(service))
	r.Get("/{id}", getResourceHandler(service))
	r.Get("/", getResourcesHandler(service))
}
