package service

import (
	"context"
	"path"

	model2 "miniWiki/internal/domain/image/model"
	"miniWiki/internal/domain/resource/model"

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

	err = s.resourceRepository.UpdateResourcePicture(ctx,
		request.ResourceId,
		request.AccountId,
		filename,
	)
	if err != nil {
		logrus.WithContext(ctx).
			WithField("resource_id", request.ResourceId).
			Info(err)
		return err
	}

	return err
}
