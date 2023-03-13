package service

import (
	"context"

	"miniWiki/domain/profile/model"

	"github.com/sirupsen/logrus"
)

func (s *Profile) CreateProfile(ctx context.Context, request model.CreateProfileRequest) error {
	_, err := s.profileQuerier.CreateProfile(ctx, request.AccountId, request.Profile.Name)
	if err != nil {
		logrus.WithContext(ctx).
			WithField("account_id", request.AccountId).
			Infof("Failed creating profile %v", err)
		return err
	}

	if request.Profile.Alias != nil {
		_, err = s.profileQuerier.UpdateAlias(ctx, *request.Profile.Alias, request.AccountId)
		if err != nil {
			logrus.WithContext(ctx).
				WithField("account_id", request.AccountId).
				Infof("Failed adding alias %v", err)
			return err
		}
	}

	return nil
}
