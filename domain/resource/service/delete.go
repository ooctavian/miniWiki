package service

import (
	"context"

	"miniWiki/domain/resource/model"

	"github.com/sirupsen/logrus"
)

func (s *Resource) DeleteResource(ctx context.Context, request model.DeleteResourceRequest) error {
	_, err := s.resourceQuerier.DeleteResourceByID(ctx, request.ResourceId, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).
			WithField("resource_id", request.ResourceId).
			Errorf("Failed deleting from database: %v", err)
	}

	return err
}
