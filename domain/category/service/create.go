package service

import (
	"context"

	"miniWiki/domain/category/model"

	"github.com/sirupsen/logrus"
)

func (s *Category) CreateCategory(ctx context.Context, request model.CreateCategoryRequest) (*int, error) {
	err := s.db.
		Create(&request.Category).
		Error
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed inserting category in database: %v", err)
		return nil, err
	}

	return &request.Category.ID, nil
}
