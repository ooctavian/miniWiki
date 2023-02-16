package service

import (
	"context"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"
)

func (s *Resource) DeleteResource(ctx context.Context, request model.DeleteResourceRequest) error {
	_, err := s.resourceQuerier.DeleteResourceByID(ctx, request.ResourceId)
	if err != nil {
		utils.Logger.WithContext(ctx).Errorf("Failed deleting from database: %v", err)
	}

	return err
}
