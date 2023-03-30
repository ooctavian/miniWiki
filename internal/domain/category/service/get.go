package service

import (
	"context"

	"miniWiki/pkg/utils"

	"github.com/sirupsen/logrus"
)

func (s *Category) GetCategories(ctx context.Context, pagination utils.Pagination) (*utils.Pagination, error) {
	categories, err := s.categoryRepository.GetCategories(ctx, pagination)
	if err != nil {
		logrus.WithContext(ctx).Infof("Error getting categories: %v", err)
		return nil, err
	}
	return &categories, nil
}
