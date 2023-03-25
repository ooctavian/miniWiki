package service

import (
	"context"
	"path"

	model2 "miniWiki/pkg/domain/image/model"
	"miniWiki/pkg/domain/resource/model"
	"miniWiki/pkg/domain/resource/query"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *Resource) UploadResourceImage(ctx context.Context, request model.UploadResourceImageRequest) error {
	filename := uuid.NewString() + path.Ext(request.ImageName)
	req := model2.UploadRequest{
		ImageFolder: "resources",
		ImageName:   filename,
		Image:       request.Image,
	}
	err := s.imageService.Upload(ctx, req)
	if err != nil {
		logrus.
			WithContext(ctx).
			WithField("resource_id", request.ResourceId).
			Info("Error uploading file", err)
		return err
	}

	params := query.UpdateResourceImageParams{
		ImageUrl:   filename,
		ResourceID: request.ResourceId,
		AuthorID:   request.AccountId,
	}
	_, err = s.resourceQuerier.UpdateResourceImage(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).
			WithField("resource_id", request.ResourceId).
			Info(err)
		return err
	}

	return err
}
