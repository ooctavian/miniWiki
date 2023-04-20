package service

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) GetAccount(ctx context.Context, request model.GetAccountRequest) (*model.AccountResponse, error) {
	acc, err := s.accountRepository.GetAccount(ctx, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed getting account: %v", err)
		return nil, err
	}

	return &model.AccountResponse{
		Email:      acc.Email,
		Active:     acc.Active,
		Name:       acc.Name,
		Alias:      acc.Alias,
		PictureUrl: acc.PictureUrl,
		CreatedAt:  acc.CreatedAt,
		UpdatedAt:  acc.UpdatedAt,
	}, nil
}
