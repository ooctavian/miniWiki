package service

import (
	"context"
	"strings"

	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"

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

	res, err := s.resourceQuerier.InsertResource(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).
			Errorf("Failed inserting in database: %v", err)
	}

	return &model.ResourceResponse{
		ResourceId: res,
	}, err
}
