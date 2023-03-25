package service

import (
	"context"
	"path"

	"miniWiki/pkg/domain/account/model"
	model2 "miniWiki/pkg/domain/image/model"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *Account) UploadProfilePicture(ctx context.Context, request model.UploadProfilePictureRequest) error {
	filename := uuid.NewString() + path.Ext(request.ImageName)
	req := model2.UploadRequest{
		ImageFolder: "account",
		ImageName:   filename,
		Image:       request.Image,
	}

	err := s.imageService.Upload(ctx, req)
	if err != nil {
		logrus.WithContext(ctx).
			WithField("account_id", request.AccountId).
			Info("Error uploading file", err)
		return err
	}

	err = s.db.
		Model(&model.Account{}).
		Where("account_id = ?", request.AccountId).
		Update("picture_url", filename).
		Error
	if err != nil {
		logrus.WithContext(ctx).WithField("account_id", request.AccountId).Info(err)
		return err
	}

	return err
}
