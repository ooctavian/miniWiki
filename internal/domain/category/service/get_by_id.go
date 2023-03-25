package service

import (
	"context"
	"errors"
	"strconv"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Category) GetCategory(ctx context.Context, request model.GetCategoryRequest) (*model.Category, error) {
	category, err := s.categoryRepository.GetCategory(ctx, request.CategoryId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, transport.NotFoundError{
				Item: "category",
				Id:   strconv.Itoa(request.CategoryId),
			}
		}
		logrus.WithContext(ctx).Infof("Error getting category by id: %v", err)
		return nil, err
	}

	return category, nil
}
