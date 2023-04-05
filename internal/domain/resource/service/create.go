package service

import (
	"context"

	"miniWiki/internal/domain/resource/model"

	"github.com/sirupsen/logrus"
)

func (s *Resource) CreateResource(ctx context.Context, request model.CreateResourceRequest) (*uint, error) {
	request.Resource.AuthorId = request.AccountId
	id, err := s.resourceRepository.InsertResource(ctx, request.Resource)
	if err != nil {
		logrus.WithContext(ctx).
			Errorf("Failed inserting in database: %v", err)
		return nil, err
	}

	return &id, err
}
