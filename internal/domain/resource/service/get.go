package service

import (
	"context"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"

	"github.com/sirupsen/logrus"
)

func (s *Resource) GetResources(ctx context.Context, request model.GetResourcesRequest) (utils.Pagination, error) {
	resources, err := s.resourceRepository.GetResources(ctx, request.AccountId, request.Pagination, request.Filters)

	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed retrieving resources: %v", err)
		return request.Pagination, err
	}

	return resources, nil
}
