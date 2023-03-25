package service

import (
	"context"

	"miniWiki/pkg/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) GetPublicAccount(ctx context.Context,
	request model.GetAccountRequest) (*model.PublicAccountResponse, error) {
	var account model.PublicAccountResponse
	err := s.db.Model(&model.Account{}).First(&account, request.AccountId).Error
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed getting account: %v", err)
		return nil, err
	}

	return &account, nil
}
