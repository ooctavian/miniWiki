package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"

	"github.com/sirupsen/logrus"
)

func (s *Resource) GetResources(ctx context.Context, request model.GetResourcesRequest) ([]model.ResourceResponse, error) {
	getResourcesRow, err := s.resourceQuerier.GetResources(ctx,
		query.GetResourcesParams{
			Title:      ptrToString(request.Filters.Title),
			Link:       ptrToString(request.Filters.Link),
			Categories: request.Filters.Categories,
		},
	)

	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed retrieving resources: %v", err)
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
				CategoryId:  r.CategoryID,
			})
	}

	return resources, nil
}

func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
