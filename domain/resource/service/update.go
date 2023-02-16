package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"
	"miniWiki/utils"
)

func (s *Resource) UpdateResource(ctx context.Context, request model.UpdateResourceRequest) error {
	params := query.UpdateResourceParams{
		ResourceID:  request.ResourceId,
		Title:       request.Resource.Title,
		Description: request.Resource.Description,
		Link:        request.Resource.Link,
	}

	_, err := s.resourceQuerier.UpdateResource(ctx, params)
	if err != nil {
		utils.Logger.WithContext(ctx).Infof("Failed updating in database: %v", err)
	}

	return err
}
