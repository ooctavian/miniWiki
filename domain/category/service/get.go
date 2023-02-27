package service

import (
	"context"

	"miniWiki/domain/category/model"

	"github.com/sirupsen/logrus"
)

func (s *Category) GetCategories(ctx context.Context) ([]model.CategoryResponse, error) {
	getCategories, err := s.categoryQuerier.GetCategories(ctx)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed inserting in database: %v", err)
		return nil, err
	}

	if len(getCategories) < 1 {
		return []model.CategoryResponse{}, nil
	}

	var response []model.CategoryResponse
	for _, c := range getCategories {
		response = append(response,
			model.CategoryResponse{
				CategoryId: *c.CategoryID,
				Title:      *c.Title,
				ParentId:   c.ParentID,
			},
		)
	}

	return response, nil
}
