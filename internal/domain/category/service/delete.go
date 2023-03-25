package service

import (
	"context"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
)

func (s *Category) DeleteCategory(ctx context.Context, request model.DeleteCategoryRequest) error {
	count, err := s.resourceRepository.CountCategoryResources(ctx, request.CategoryId)
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return err
	}
	if count > 0 {
		return transport.ForbiddenError{}
	}

	count, err = s.categoryRepository.CountCategories(ctx, request.CategoryId)
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return err
	}
	if count > 0 {
		return transport.ForbiddenError{}
	}

	err = s.categoryRepository.DeleteCategory(ctx, request.CategoryId)
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return err
	}

	return nil
}
