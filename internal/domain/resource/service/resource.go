package service

import (
	cController "miniWiki/internal/domain/category/controller"
	iService "miniWiki/internal/domain/image/service"
	rQuery "miniWiki/internal/domain/resource/query"
)

type Resource struct {
	resourceQuerier rQuery.Querier
	imageService    iService.ImageService
	categoryService cController.CategoryService
}

func NewResource(rQuerier rQuery.Querier, categoryService cController.CategoryService, imageService iService.ImageService) *Resource {
	resource := &Resource{
		resourceQuerier: rQuerier,
		imageService:    imageService,
		categoryService: categoryService,
	}
	return resource
}
