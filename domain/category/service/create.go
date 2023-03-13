package service

import (
	"context"

	"miniWiki/domain/category/model"
	"miniWiki/domain/category/query"

	"github.com/sirupsen/logrus"
)

func (s *Category) CreateCategory(ctx context.Context, request model.CreateCategoryRequest) error {
	params := query.InsertCategoryParams{
		Title:    request.Category.Title,
		ParentID: int32(request.Category.ParentId),
		AuthorID: request.AccountId,
	}
	_, err := s.categoryQuerier.InsertCategory(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed inserting category in database: %v", err)
		return err
	}

	return nil

}
