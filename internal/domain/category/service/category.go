package service

import (
	cRepository "miniWiki/internal/domain/category/repository"
	rRepository "miniWiki/internal/domain/resource/repository"
)

type Category struct {
	categoryRepository cRepository.CategoryRepositoryInterface
	resourceRepository rRepository.ResourceRepositoryInterface
}

func NewCategory(
	categoryRepository cRepository.CategoryRepositoryInterface,
	resourceRepository rRepository.ResourceRepositoryInterface,
) *Category {
	category := &Category{
		categoryRepository: categoryRepository,
		resourceRepository: resourceRepository,
	}
	return category
}
