package service

import (
	"context"
	"errors"
	"io"
	"strconv"

	"miniWiki/pkg/domain/account/model"
	model2 "miniWiki/pkg/domain/image/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Account) DownloadResourceImage(ctx context.Context,
	request model.DownloadProfilePictureRequest) (io.Reader, error) {
	var account model.Account
	err := s.db.First(&account, request.AccountId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, transport.NotFoundError{
				Item: "account",
				Id:   strconv.Itoa(request.AccountId),
			}
		}
		logrus.WithContext(ctx).Infof("Error getting account by id: %v", err)
		return nil, err
	}

	req := model2.DownloadRequest{
		ImageFolder: "profile",
		ImageName:   *account.PictureUrl,
	}

	return s.imageService.Download(ctx, req)
}
