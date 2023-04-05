package service

import (
	"context"

	cController "miniWiki/internal/domain/category/controller"
	iService "miniWiki/internal/domain/image/service"
	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"
)

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
	categoryService    cController.CategoryService
}

func NewResource(
	resourceRepository resourceRepositoryInterface,
	categoryService cController.CategoryService,
	imageService iService.ImageService,
) *Resource {
	resource := &Resource{
		resourceRepository: resourceRepository,
		imageService:       imageService,
		categoryService:    categoryService,
	}
	return resource
}
