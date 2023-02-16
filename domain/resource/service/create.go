package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"
	"miniWiki/utils"
)

func (s *Resource) CreateResource(ctx context.Context, request model.CreateResourceRequest) error {
	params := query.InsertResourceParams{
		Title:       request.Resource.Title,
		Description: request.Resource.Description,
		Link:        request.Resource.Link,
	}

	_, err := s.resourceQuerier.InsertResource(ctx, params)
	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed inserting in database: %v", err)
	}
	return err
}
