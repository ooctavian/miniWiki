package service

import (
	"context"
	"errors"
	"io"
	"strconv"

	model2 "miniWiki/domain/image/model"
	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"
	"miniWiki/utils"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Resource) DownloadResourceImage(ctx context.Context, request model.DownloadResourceImageRequest) (io.Reader, error) {
	resource, err := s.getResource(ctx, request.ResourceId, request.AccountId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Resource not found: %v", err)
			return nil, utils.NotFoundError{
				Item: "resource",
				Id:   strconv.Itoa(request.ResourceId),
			}
		}
		logrus.WithContext(ctx).Errorf("Failed retrieving resource: %v", err)
		return nil, err
	}

	if resource.State == query.ResourceStatePRIVATE &&
		*resource.AuthorID != request.AccountId {
		return nil, utils.ForbiddenError{}
	}

	if resource.Image == nil {
		return nil, utils.NotFoundError{
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
