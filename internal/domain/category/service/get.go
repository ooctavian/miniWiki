package service

import (
	"context"

	"miniWiki/internal/domain/category/model"

	"github.com/sirupsen/logrus"
)

func (s *Category) GetCategories(ctx context.Context) ([]model.Category, error) {
	categories, err := s.categoryRepository.GetCategories(ctx)
	if err != nil {
		logrus.WithContext(ctx).Infof("Error getting categories: %v", err)
		return nil, err
	}

	return categories, nil
}
