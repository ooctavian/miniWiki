package service

import (
	"context"
	"strings"

	model2 "miniWiki/internal/domain/category/model"
	"miniWiki/internal/domain/resource/model"
	"miniWiki/internal/domain/resource/query"

	"github.com/sirupsen/logrus"
)

func (s *Resource) CreateResource(ctx context.Context,
	request model.CreateResourceRequest) (*model.ResourceResponse, error) {
	params := query.InsertResourceParams{
		Title:       request.Resource.Title,
		Description: request.Resource.Description,
		Link:        request.Resource.Link,
		CategoryID:  request.Resource.CategoryId,
		AuthorID:    request.AccountId,
		State:       query.ResourceState(strings.ToUpper(request.Resource.State)),
	}

	if request.Resource.CategoryName != nil {
		id, err := s.categoryService.CreateCategory(ctx, model2.CreateCategoryRequest{
			Category: model2.CreateCategory{
				Title: *request.Resource.CategoryName,
			},
		})
		if err != nil {
			logrus.WithContext(ctx).
				Errorf("Failed creating category: %v", err)
			return nil, err
		}
		params.CategoryID = *id
	}

	res, err := s.resourceQuerier.InsertResource(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).
			Errorf("Failed inserting in database: %v", err)
		return nil, err
	}

	return &model.ResourceResponse{
		ResourceId: res,
	}, err
}
