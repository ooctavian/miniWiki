package service

import (
	"context"

	model2 "miniWiki/internal/domain/category/model"
	"miniWiki/internal/domain/resource/model"

	"github.com/sirupsen/logrus"
)

func (s *Resource) CreateResource(ctx context.Context, request model.CreateResourceRequest) (*uint, error) {
	request.Resource.AuthorId = request.AccountId
	if request.Resource.CategoryName != nil {
		id, err := s.categoryService.CreateCategory(ctx, model2.CreateCategoryRequest{
			Category: model2.CreateCategory{
				Title: *request.Resource.CategoryName,
			},
		})

		if err != nil {
			logrus.WithContext(ctx).
				Errorf("Failed creating category: %v", err)
			return nil, err
		}
		request.Resource.CategoryId = *id
	}

	id, err := s.resourceRepository.InsertResource(ctx, request.Resource)
	if err != nil {
		logrus.WithContext(ctx).
			Errorf("Failed inserting in database: %v", err)
		return nil, err
	}

	return &id, err
}
