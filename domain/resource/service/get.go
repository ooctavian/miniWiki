package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"
)

func (s *Resource) GetResources(ctx context.Context) ([]model.ResourceResponse, error) {
	getResourcesRow, err := s.resourceQuerier.GetResources(ctx)

	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed retrieving resources: %v", err)
		return nil, err
	}

	var resources []model.ResourceResponse
	for _, r := range getResourcesRow {
		resources = append(resources,
			model.ResourceResponse{
				ResourceId:  *r.ResourceID,
				Title:       *r.Title,
				Description: *r.Description,
				Link:        *r.Link,
			})
	}

	return resources, nil
}
