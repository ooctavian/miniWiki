package service

import (
	"context"
	"errors"
	"strconv"

	"miniWiki/pkg/domain/resource/model"
	"miniWiki/pkg/domain/resource/query"
	"miniWiki/pkg/transport"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Resource) GetResource(ctx context.Context, request model.GetResourceRequest) (*model.ResourceResponse, error) {
	resource, err := s.getResource(ctx, request.ResourceId, request.AccountId)
	if err != nil {
		return nil, err
	}

	response := &model.ResourceResponse{
		ResourceId:  resource.ResourceID,
		Title:       *resource.Title,
		Description: *resource.Description,
		Link:        resource.Link,
		State:       string(resource.State),
		CategoryId:  resource.CategoryID,
		AuthorId:    *resource.AuthorID,
	}

	return response, nil
}

func (s *Resource) getResource(ctx context.Context, resourceId int, accountId int) (*query.GetResourceByIDRow, error) {
	resource, err := s.resourceQuerier.GetResourceByID(ctx, resourceId)

	//NOTE: Should it be extracted in another function ?
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Resource not found: %v", err)
			return nil, transport.NotFoundError{
				Item: "resource",
				Id:   strconv.Itoa(resourceId),
			}
		}
		logrus.WithContext(ctx).Errorf("Failed retrieving resource: %v", err)
		return nil, err
	}

	if resource.State == query.ResourceStatePRIVATE &&
		*resource.AuthorID != accountId {
		return nil, transport.ForbiddenError{}
	}

	return &resource, nil
}
