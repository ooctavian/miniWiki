package service

import (
	"context"

	"miniWiki/domain/category/model"
	"miniWiki/domain/category/query"

	"github.com/sirupsen/logrus"
)

func (s *Category) UpdateCategory(ctx context.Context, request model.UpdateCategoryRequest) error {
	params := query.UpdateCategoryParams{
		Title:      request.Category.Title,
		ParentID:   int32(request.Category.ParentId),
		CategoryID: request.CategoryId,
		AuthorID:   request.AccountId,
	}

	_, err := s.categoryQuerier.UpdateCategory(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed updating category: %v", err)
	}

	return err
}
