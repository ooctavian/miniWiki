package service

import (
	"context"

	"miniWiki/internal/domain/account/model"
	"miniWiki/pkg/security"

	"github.com/sirupsen/logrus"
)

func (s *Account) UpdateAccount(ctx context.Context, request model.UpdateAccountRequest) error {
	account := request.Account
	if request.Account.Password != nil {
		err := security.ValidatePassword([]byte(*request.Account.Password))
		if err != nil {
			logrus.WithContext(ctx).Errorf("Weak password: %v", err)
			return err
		}
		encryptedPassword, err := s.hash.GenerateFormatted(*request.Account.Password)
		if err != nil {
			logrus.WithContext(ctx).Errorf("Failed generating security: %v", err)
			return err
		}
		account.Password = &encryptedPassword
	}

	err := s.accountRepository.UpdateAccount(ctx, request.AccountId, request.Account)

	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed updating account: %v", err)
		return err
	}

	return nil
}
