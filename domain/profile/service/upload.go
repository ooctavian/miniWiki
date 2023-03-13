package service

import (
	"context"
	"path"

	model2 "miniWiki/domain/image/model"
	"miniWiki/domain/profile/model"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *Profile) UploadProfilePicture(ctx context.Context, request model.UploadProfilePictureRequest) error {
	filename := uuid.NewString() + path.Ext(request.ImageName)
	req := model2.UploadRequest{
		ImageFolder: "profile",
		ImageName:   filename,
		Image:       request.Image,
	}

	err := s.imageService.Upload(ctx, req)
	if err != nil {
		logrus.WithContext(ctx).Info("Error uploading file", err)
		return err
	}

	_, err = s.profileQuerier.UpdateProfilePicture(ctx, filename, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).WithField("account_id", request.AccountId).Info(err)
		return err
	}

	return err
}
