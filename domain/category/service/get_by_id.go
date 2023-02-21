package service

import (
	"context"
	"errors"

	"miniWiki/domain/category/model"
	"miniWiki/utils"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Category) GetCategory(ctx context.Context, request model.GetCategoryRequest) (*model.CategoryResponse, error) {
	getCategory, err := s.categoryQuerier.GetCategoryByID(ctx, request.CategoryId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Category not found: %v", err)
			return nil, &utils.NotFoundError{
				Item: "category",
				Id:   string(request.CategoryId),
			}
		}
		logrus.WithContext(ctx).Errorf("Failed retrieving category: %v", err)
		return nil, err
	}

	resource := &model.CategoryResponse{
		CategoryId: getCategory.CategoryID,
		Title:      getCategory.Title,
		ParentId:   getCategory.ParentID,
	}

	return resource, nil
}
