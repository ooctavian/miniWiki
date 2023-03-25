package service

import (
	"context"

	"miniWiki/pkg/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) UpdateAccount(ctx context.Context, request model.UpdateAccountRequest) error {
	account := request.Account
	if request.Account.Password != nil {
		encryptedPassword, err := s.hash.GenerateFormatted(*request.Account.Password)
		if err != nil {
			logrus.WithContext(ctx).Errorf("Failed generating security: %v", err)
			return err
		}
		account.Password = &encryptedPassword
	}

	err := s.db.Model(&model.Account{}).
		Where("account_id = ?", request.AccountId).
		Updates(account).
		Error

	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed updating account: %v", err)
		return err
	}

	return nil
}
