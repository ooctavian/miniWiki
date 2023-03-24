package service

import (
	"context"
	"errors"
	"strconv"

	"miniWiki/domain/category/model"
	"miniWiki/transport"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Category) GetCategory(ctx context.Context, request model.GetCategoryRequest) (*model.Category, error) {
	var category model.Category
	err := s.db.First(&category, request.CategoryId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, transport.NotFoundError{
				Item: "category",
				Id:   strconv.Itoa(request.CategoryId),
			}
		}
		logrus.WithContext(ctx).Info("Error getting category by id: %v", err)
		return nil, err
	}

	return &category, nil
}
