package service

import (
	"context"

	model2 "miniWiki/internal/domain/category/model"
	"miniWiki/internal/domain/resource/model"

	"github.com/sirupsen/logrus"
)

func (s *Resource) UpdateResource(ctx context.Context, request model.UpdateResourceRequest) error {
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
			return err
		}
		request.Resource.CategoryId = id
	}

	err := s.resourceRepository.UpdateResource(ctx, request)
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed updating in database: %v", err)
		return err
	}

	return err
}
