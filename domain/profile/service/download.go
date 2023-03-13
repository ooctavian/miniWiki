package service

import (
	"context"
	"errors"
	"io"
	"strconv"

	model2 "miniWiki/domain/image/model"
	"miniWiki/domain/profile/model"
	"miniWiki/utils"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Profile) DownloadResourceImage(ctx context.Context, request model.DownloadProfilePictureRequest) (io.Reader, error) {
	profile, err := s.profileQuerier.GetProfile(ctx, request.AccountId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Profile not found: %v", err)
			return nil, utils.NotFoundError{
				Item: "profile",
				Id:   strconv.Itoa(request.AccountId),
			}
		}

		logrus.WithContext(ctx).Errorf("Failed retrieving resource: %v", err)
		return nil, err
	}

	req := model2.DownloadRequest{
		ImageFolder: "profile",
		ImageName:   *profile.PictureUrl,
	}

	return s.imageService.Download(ctx, req)
}
