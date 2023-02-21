package service

import (
	"context"

	"miniWiki/domain/category/model"

	"github.com/sirupsen/logrus"
)

func (s *Category) UpdateCategory(ctx context.Context, request model.UpdateCategoryRequest) error {
	var err error
	if request.Category.ParentId == nil {
		_, err = s.categoryQuerier.UpdateCategory(ctx, *request.Category.Title, request.CategoryId)
		if err != nil {
			logrus.WithContext(ctx).Errorf("Failed updating category: %v", err)
		}
	}

	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed updating database: %v", err)
	}

	return err
}
