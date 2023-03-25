package service

import (
	"context"
	"errors"
	"io"
	"strconv"

	"miniWiki/internal/domain/account/model"
	model2 "miniWiki/internal/domain/image/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (s *Account) DownloadResourceImage(ctx context.Context,
	request model.DownloadProfilePictureRequest) (io.Reader, error) {
	acc, err := s.accountRepository.GetAccount(ctx, request.AccountId)
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
		ImageName:   *acc.PictureUrl,
	}

	return s.imageService.Download(ctx, req)
}
