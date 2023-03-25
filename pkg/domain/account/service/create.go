package service

import (
	"context"

	"miniWiki/pkg/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) CreateAccount(ctx context.Context, request model.CreateAccountRequest) error {
	account := request.Account
	encryptedPassword, err := s.hash.GenerateFormatted(account.Password)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed generating security: %v", err)
		return err
	}

	account.Password = encryptedPassword
	err = s.db.WithContext(ctx).Create(&account).Error
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed creating account: %v", err)
		return err
	}

	return nil
}
