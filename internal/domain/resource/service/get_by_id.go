package service

import (
	"context"
	"errors"
	"strconv"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Resource) GetResource(ctx context.Context, request model.GetResourceRequest) (*model.Resource, error) {
	resource, err := s.resourceRepository.GetResourceById(ctx, request.ResourceId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithContext(ctx).Errorf("Resource not found: %v", err)
			return nil, transport.NotFoundError{
				Item: "resource",
				Id:   strconv.Itoa(request.ResourceId),
			}
		}
		logrus.WithContext(ctx).Errorf("Failed retrieving resource: %v", err)
		return nil, err
	}

	if resource.State == "PRIVATE" &&
		resource.AuthorId != uint(request.AccountId) {
		return nil, transport.ForbiddenError{}
	}

	return resource, nil
}
