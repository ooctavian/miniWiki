package service

import (
	"context"

	"miniWiki/domain/category/model"
	rmodel "miniWiki/domain/resource/model"
	"miniWiki/transport"

	"github.com/sirupsen/logrus"
)

func (s *Category) DeleteCategory(ctx context.Context, request model.DeleteCategoryRequest) error {
	var count int64
	err := s.db.Model(&rmodel.Resource{}).
		Where("category_id = ?", request.CategoryId).
		Count(&count).Error
	if err != nil {
		logrus.WithContext(ctx).Info("Failed deleting category: %v", err)
		return nil
	}
	if count > 0 {
		return transport.ForbiddenError{}
	}

	err = s.db.Model(&model.Category{}).
		Where("parent_id = ?", request.CategoryId).
		Count(&count).Error
	if err != nil {
		logrus.WithContext(ctx).Info("Failed deleting category: %v", err)
		return nil
	}
	if count > 0 {
		return transport.ForbiddenError{}
	}

	var category model.Category
	err = s.db.Delete(category, request.CategoryId).Error
	if err != nil {
		logrus.WithContext(ctx).Info("Failed deleting category: %v", err)
		return nil
	}

	return nil
}
