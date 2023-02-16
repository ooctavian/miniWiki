package service

import (
	"context"

	"miniWiki/domain/category/model"
	"miniWiki/utils"
)

func (s *Category) DeleteCategory(ctx context.Context, request model.DeleteCategoryRequest) error {
	_, err := s.categoryQuerier.DeleteCategoryByID(ctx, request.CategoryId)

	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed deleting category: %v", err)
		return err
	}

	return nil
}
