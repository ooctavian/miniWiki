package service

import (
	"context"

	model2 "miniWiki/internal/domain/category/model"
	iService "miniWiki/internal/domain/image/service"
	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"
)

type categoryServiceInterface interface {
	CreateCategory(ctx context.Context, request model2.CreateCategoryRequest) (*int, error)
}

type resourceRepositoryInterface interface {
	GetResourceById(ctx context.Context, id int) (*model.Resource, error)
	GetResources(ctx context.Context, accountId int, pagination utils.Pagination, filters model.GetResourcesFilters) (utils.Pagination, error)
	DeleteResourceById(ctx context.Context, resourceId uint, accountId uint) error
	InsertResource(ctx context.Context, resource model.CreateResource) (uint, error)
	UpdateResource(ctx context.Context, request model.UpdateResourceRequest) error
	UpdateResourcePicture(ctx context.Context, resourceId int, accountId int, path string) error
}

type Resource struct {
	resourceRepository resourceRepositoryInterface
	imageService       iService.ImageService
	categoryService    categoryServiceInterface
}

func NewResource(
	resourceRepository resourceRepositoryInterface,
	categoryService categoryServiceInterface,
	imageService iService.ImageService,
) *Resource {
	resource := &Resource{
		resourceRepository: resourceRepository,
		imageService:       imageService,
		categoryService:    categoryService,
	}
	return resource
}
