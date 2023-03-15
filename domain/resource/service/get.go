package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"

	"github.com/sirupsen/logrus"
)

func (s *Resource) GetResources(ctx context.Context, request model.GetResourcesRequest) ([]model.ResourceResponse, error) {
	resources, err := s.resourceQuerier.GetResources(ctx,
		query.GetResourcesParams{
			Title:      request.Filters.Title,
			Link:       request.Filters.Link,
			Categories: request.Filters.Categories,
			AccountID:  request.AccountId,
		},
	)

	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed retrieving resources: %v", err)
		return nil, err
	}

	if len(resources) < 1 {
		return []model.ResourceResponse{}, nil
	}

	var response []model.ResourceResponse

	for _, r := range resources {
		response = append(response,
			model.ResourceResponse{
				ResourceId:  *r.ResourceID,
				Title:       *r.Title,
				Description: *r.Description,
				Link:        *r.Link,
				CategoryId:  r.CategoryID,
				State:       string(r.State),
				AuthorId:    *r.AuthorID,
			})
	}

	return response, nil
}
