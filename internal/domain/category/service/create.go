package service

import (
	"context"

	"miniWiki/internal/domain/category/model"

	"github.com/sirupsen/logrus"
)

func (s *Category) CreateCategory(ctx context.Context, request model.CreateCategoryRequest) (*int, error) {
	categoryId, err := s.categoryRepository.CreateCategory(ctx, request.Category)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed inserting category in database: %v", err)
		return nil, err
	}

	return &categoryId, nil
}
