package service

import (
	"context"

	"miniWiki/domain/category/model"
	"miniWiki/utils"
)

func (s *Category) UpdateCategory(ctx context.Context, request model.UpdateCategoryRequest) error {
	var err error
	if request.Category.ParentId == nil {
		_, err = s.categoryQuerier.UpdateCategory(ctx, *request.Category.Title, request.CategoryId)
		if err != nil {
			utils.Logger.WithContext(ctx).Errorf("Failed updating category: %v", err)
		}
	}

	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed updating database: %v", err)
	}

	return err
}
