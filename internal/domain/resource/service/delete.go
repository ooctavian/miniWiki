package service

import (
	"context"

	"miniWiki/internal/domain/resource/model"

	"github.com/sirupsen/logrus"
)

func (s *Resource) DeleteResource(ctx context.Context, request model.DeleteResourceRequest) error {
	err := s.resourceRepository.DeleteResourceById(ctx, uint(request.ResourceId), uint(request.AccountId))
	if err != nil {
		logrus.WithContext(ctx).
			WithField("resource_id", request.ResourceId).
			Errorf("Failed deleting from database: %v", err)
	}

	return err
}
