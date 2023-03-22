package service

import (
	"context"

	"miniWiki/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) DeactivateAccount(ctx context.Context, request model.DeactivateAccountRequest) error {
	_, err := s.resourceQuerier.MakeAccountResourcesPrivate(ctx, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed deactivating account: %v", err)
		return err
	}

	_, err = s.accountQuerier.UpdateAccountStatus(ctx, false, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed changing account status: %v", err)
		return err
	}
	return nil
}
