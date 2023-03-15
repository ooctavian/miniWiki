package service

import (
	"context"
	"errors"
	"strconv"

	"miniWiki/domain/profile/model"
	"miniWiki/utils"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Profile) GetProfile(ctx context.Context, request model.GetProfileRequest) (*model.ProfileResponse, error) {
	profile, err := s.profileQuerier.GetProfile(ctx, request.AccountId)

	//NOTE: Should it be extracted in another function ?
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).
				WithField("account_id", profile.AccountID).
				Errorf("Profile not found: %v", err)
			return nil, utils.NotFoundError{
				Item: "profile",
				Id:   strconv.Itoa(request.AccountId),
			}
		}
		logrus.WithContext(ctx).Errorf("Failed retrieving profile: %v", err)
		return nil, err
	}

	response := &model.ProfileResponse{
		Name:       profile.Name,
		Alias:      profile.Alias,
		PictureUrl: profile.PictureUrl,
	}

	return response, nil
}
