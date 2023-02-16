package service

import (
	"context"

	"miniWiki/domain/category/model"
	"miniWiki/utils"
)

func (s *Category) GetCategory(ctx context.Context, request model.GetCategoryRequest) (*model.CategoryResponse, error) {
	getCategory, err := s.categoryQuerier.GetCategoryByID(ctx, request.CategoryId)

	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed retrieving category: %v", err)
		return nil, err
	}

	resource := &model.CategoryResponse{
		CategoryId: getCategory.CategoryID,
		Title:      getCategory.Title,
		ParentId:   getCategory.ParentID,
	}

	return resource, nil
}
