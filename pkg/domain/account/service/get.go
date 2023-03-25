package service

import (
	"context"

	"miniWiki/pkg/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) GetAccount(ctx context.Context, request model.GetAccountRequest) (*model.AccountResponse, error) {
	var account model.AccountResponse
	err := s.db.Model(&model.Account{}).First(&account, request.AccountId).Error
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed getting account: %v", err)
		return nil, err
	}

	return &account, nil
}
