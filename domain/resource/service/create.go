package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"

	"github.com/sirupsen/logrus"
)

func (s *Resource) CreateResource(ctx context.Context, request model.CreateResourceRequest) error {
	params := query.InsertResourceParams{
		Title:       request.Resource.Title,
		Description: request.Resource.Description,
		Link:        request.Resource.Link,
		CategoryID:  request.Resource.CategoryId,
	}

	_, err := s.resourceQuerier.InsertResource(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed inserting in database: %v", err)
	}

	return err
}
