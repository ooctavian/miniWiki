package service

import (
	"context"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/utils"
)

type CategoryRepositoryInterface interface {
	CreateCategory(ctx context.Context, category model.CreateCategory) (int, error)
	GetCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error)
	GetCategory(ctx context.Context, id int) (model.Category, error)
	DeleteCategory(ctx context.Context, id int) error
	CountCategories(ctx context.Context, id int) (int64, error)
}

type ResourceRepositoryInterface interface {
	CountCategoryResources(ctx context.Context, id int) (int64, error)
}

type Category struct {
	categoryRepository CategoryRepositoryInterface
	resourceRepository ResourceRepositoryInterface
}

func NewCategory(
	categoryRepository CategoryRepositoryInterface,
	resourceRepository ResourceRepositoryInterface,
) *Category {
	category := &Category{
		categoryRepository: categoryRepository,
		resourceRepository: resourceRepository,
	}
	return category
}
