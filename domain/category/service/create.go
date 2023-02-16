package service

import (
	"context"

	"miniWiki/domain/category/model"
	"miniWiki/utils"
)

func (s *Category) CreateCategory(ctx context.Context, request model.CreateCategoryRequest) error {
	var err error
	if request.Category.ParentId == nil {
		_, err = s.categoryQuerier.InsertCategory(ctx, request.Category.Title)
		if err != nil {
			utils.Logger.WithContext(ctx).Errorf("Failed inserting category in database: %v", err)
			return err
		}

		return nil
	}
	_, err = s.categoryQuerier.InsertSubCategory(ctx, request.Category.Title, *request.Category.ParentId)
	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed inserting subcategory in database: %v", err)
	}

	return nil
}
