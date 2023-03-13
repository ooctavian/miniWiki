package service

import (
	"context"

	"miniWiki/domain/profile/model"

	"github.com/sirupsen/logrus"
)

func (s *Profile) UpdateProfile(ctx context.Context, request model.UpdateProfileRequest) error {
	if request.Profile.Name != nil {
		_, err := s.profileQuerier.UpdateName(ctx, *request.Profile.Name, request.AccountId)
		if err != nil {
			logrus.WithContext(ctx).
				WithField("account_id", request.AccountId).
				Infof("Failed updating name %v", err)
			return err
		}
	}

	if request.Profile.Alias != nil {
		_, err := s.profileQuerier.UpdateAlias(ctx, *request.Profile.Alias, request.AccountId)
		if err != nil {
			logrus.WithContext(ctx).
				WithField("account_id", request.AccountId).
				Infof("Failed adding alias %v", err)
			return err
		}
	}

	return nil
}
