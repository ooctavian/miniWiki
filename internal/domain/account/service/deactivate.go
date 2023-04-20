package service

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) DeactivateAccount(ctx context.Context, request model.DeactivateAccountRequest) error {
	err := s.resourceRepository.MakeResourcesPrivate(ctx, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed deactivating account: %v", err)
		return err
	}

	inactive := false
	err = s.accountRepository.UpdateAccount(ctx, request.AccountId, model.UpdateAccount{Active: &inactive})
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed changing account status: %v", err)
		return err
	}
	return nil
}
