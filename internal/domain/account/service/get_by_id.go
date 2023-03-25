package service

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) GetPublicAccount(ctx context.Context,
	request model.GetAccountRequest) (*model.PublicAccountResponse, error) {
	acc, err := s.accountRepository.GetAccount(ctx, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed getting account: %v", err)
		return nil, err
	}

	return &model.PublicAccountResponse{
		Name:       acc.Name,
		Alias:      acc.Alias,
		PictureUrl: acc.PictureUrl,
	}, nil
}
