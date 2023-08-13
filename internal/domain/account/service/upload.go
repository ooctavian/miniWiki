package service

import (
	"context"
	"path"

	"miniWiki/internal/domain/account/model"
	model2 "miniWiki/internal/domain/filemanager/model"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *Account) UploadProfilePicture(ctx context.Context, request model.UploadProfilePictureRequest) error {
	filename := uuid.NewString() + path.Ext(request.ImageName)
	req := model2.UploadRequest{
		Folder:   s.imageFolder,
		Filename: filename,
		File:     request.Image,
	}

	err := s.uploader.Upload(ctx, req)
	if err != nil {
		logrus.WithContext(ctx).
			WithField("account_id", request.AccountId).
			Info("Error uploading file", err)
		return err
	}

	err = s.accountRepository.UpdateAccount(ctx, request.AccountId, model.UpdateAccount{PictureUrl: &filename})
	if err != nil {
		logrus.WithContext(ctx).WithField("account_id", request.AccountId).Info(err)
		return err
	}

	return err
}
