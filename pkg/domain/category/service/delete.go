package service

import (
	"context"

	"miniWiki/pkg/domain/category/model"
	rmodel "miniWiki/pkg/domain/resource/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
)

func (s *Category) DeleteCategory(ctx context.Context, request model.DeleteCategoryRequest) error {
	var count int64
	err := s.db.Model(&rmodel.Resource{}).
		Where("category_id = ?", request.CategoryId).
		Count(&count).Error
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return nil
	}
	if count > 0 {
		return transport.ForbiddenError{}
	}

	var category model.Category
	err = s.db.WithContext(ctx).Model(&category).
		Where("parent_id = ?", request.CategoryId).
		Count(&count).Error
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return nil
	}
	if count > 0 {
		return transport.ForbiddenError{}
	}

	err = s.db.Delete(category, request.CategoryId).Error
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return nil
	}

	return nil
}
