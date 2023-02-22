package service

import (
	"context"
	"errors"
	"fmt"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Resource) GetResource(ctx context.Context, request model.GetResourceRequest) (*model.ResourceResponse, error) {
	getResourceRow, err := s.resourceQuerier.GetResourceByID(ctx, request.ResourceId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Resource not found: %v", err)
			return nil, utils.NotFoundError{
				Item: "resource",
				Id:   fmt.Sprint(request.ResourceId),
			}
		}
		logrus.WithContext(ctx).Errorf("Failed retrieving resource: %v", err)
		return nil, err
	}

	resource := &model.ResourceResponse{
		ResourceId:  getResourceRow.ResourceID,
		Title:       *getResourceRow.Title,
		Description: *getResourceRow.Description,
		Link:        getResourceRow.Link,
		CategoryId:  getResourceRow.CategoryID,
	}

	return resource, nil
}
