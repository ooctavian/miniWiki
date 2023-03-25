package service

import (
	"context"

	"miniWiki/pkg/domain/category/model"

	"github.com/sirupsen/logrus"
)

func (s *Category) GetCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := s.db.Find(&categories).Error
	if err != nil {
		logrus.WithContext(ctx).Infof("Error getting categories: %v", err)
		return nil, err
	}

	return categories, nil
}
