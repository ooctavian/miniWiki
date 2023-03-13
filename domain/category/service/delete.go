package service

import (
	"context"

	"miniWiki/domain/category/model"

	"github.com/sirupsen/logrus"
)

func (s *Category) DeleteCategory(ctx context.Context, request model.DeleteCategoryRequest) error {
	_, err := s.categoryQuerier.DeleteCategoryByID(ctx, request.CategoryId, request.AccountId)

	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed deleting category: %v", err)
		return err
	}

	return nil
}
