package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"
)

func (s *Resource) GetResource(ctx context.Context, request model.GetResourceRequest) (*model.ResourceResponse, error) {
	getResourceRow, err := s.resourceQuerier.GetResourceByID(ctx, request.ResourceId)

	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed retrieving resource: %v", err)
		return nil, err
	}

	resource := &model.ResourceResponse{
		ResourceId:  getResourceRow.ResourceID,
		Title:       *getResourceRow.Title,
		Description: *getResourceRow.Description,
		Link:        *getResourceRow.Link,
	}

	return resource, nil
}
