package service

import (
	"context"

	"miniWiki/domain/account/model"
	"miniWiki/domain/account/query"

	"github.com/sirupsen/logrus"
)

func (s *Account) UpdateAccount(ctx context.Context, request model.UpdateAccountRequest) error {
	acc, err := s.accountQuerier.GetAccountByID(ctx, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed getting account: %v", err)
		return err
	}

	params := query.UpdateAccountParams{
		Email:     acc.Email,
		Password:  *acc.Password,
		AccountID: request.AccountId,
	}

	if request.Account.Password != nil {
		encryptedPassword, err := s.hash.GenerateFormatted(*request.Account.Password)
		if err != nil {
			logrus.WithContext(ctx).Errorf("Failed generating security: %v", err)
			return err
		}
		params.Password = encryptedPassword
	}

	if request.Account.Email != nil {
		params.Email = *request.Account.Email
	}

	_, err = s.accountQuerier.UpdateAccount(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed updating account: %v", err)
		return err
	}

	return nil
}
