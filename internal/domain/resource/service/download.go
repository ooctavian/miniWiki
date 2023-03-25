package service

import (
	"context"
	"errors"
	"io"
	"strconv"

	model2 "miniWiki/internal/domain/image/model"
	"miniWiki/internal/domain/resource/model"
	"miniWiki/internal/domain/resource/query"
	"miniWiki/pkg/transport"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Resource) DownloadResourceImage(ctx context.Context,
	request model.DownloadResourceImageRequest) (io.Reader, error) {
	resource, err := s.getResource(ctx, request.ResourceId, request.AccountId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Resource not found: %v", err)
			return nil, transport.NotFoundError{
				Item: "resource",
				Id:   strconv.Itoa(request.ResourceId),
			}
		}
		logrus.
			WithContext(ctx).
			WithField("resource_id", request.ResourceId).
			Errorf("Failed retrieving resource: %v", err)
		return nil, err
	}

	if resource.State == query.ResourceStatePRIVATE &&
		*resource.AuthorID != request.AccountId {
		return nil, transport.ForbiddenError{}
	}

	if resource.Image == nil {
		return nil, transport.NotFoundError{
			Item: "resource image",
			Id:   strconv.Itoa(request.ResourceId),
		}
	}

	req := model2.DownloadRequest{
		ImageFolder: "resources",
		ImageName:   *resource.Image,
	}

	return s.imageService.Download(ctx, req)
}
