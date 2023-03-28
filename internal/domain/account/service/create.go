package service

import (
	"context"

	"miniWiki/internal/domain/account/model"
	"miniWiki/pkg/security"

	"github.com/sirupsen/logrus"
)

func (s *Account) CreateAccount(ctx context.Context, request model.CreateAccountRequest) error {
	err := security.ValidatePassword([]byte(request.Account.Password))
	if err != nil {
		logrus.WithContext(ctx).Errorf("Weak password: %v", err)
		return err
	}

	account := request.Account
	encryptedPassword, err := s.hash.GenerateFormatted(account.Password)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed generating security: %v", err)
		return err
	}

	account.Password = encryptedPassword
	err = s.accountRepository.CreateAccount(ctx, account)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed creating account: %v", err)
		return err
	}

	return nil
}
