package service

import (
	"context"

	"miniWiki/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) CreateAccount(ctx context.Context, request model.CreateAccountRequest) error {
	encryptedPassword, err := s.hash.GenerateFormatted(request.Account.Password)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed generating security: %v", err)
		return err
	}

	_, err = s.accountQuerier.CreateAccount(ctx, request.Account.Email, encryptedPassword)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed creating account: %v", err)
		return err
	}

	return nil
}
