package service

import (
	"context"
	"errors"
	"fmt"

	"miniWiki/domain/category/model"
	"miniWiki/utils"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Category) GetCategory(ctx context.Context, request model.GetCategoryRequest) (*model.CategoryResponse, error) {
	category, err := s.categoryQuerier.GetCategoryByID(ctx, request.CategoryId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Category not found: %v", err)
			return nil, utils.NotFoundError{
				Item: "category",
				Id:   fmt.Sprint(request.CategoryId),
			}
		}
		logrus.WithContext(ctx).Errorf("Failed retrieving category: %v", err)
		return nil, err
	}

	response := &model.CategoryResponse{
		CategoryId: category.CategoryID,
		Title:      category.Title,
		ParentId:   category.ParentID,
	}

	return response, nil
}
