package service

import (
	"context"

	"miniWiki/domain/account/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) GetAccount(ctx context.Context, request model.GetAccountRequest) (*model.AccountResponse, error) {
	acc, err := s.accountQuerier.GetAccountByID(ctx, request.AccountId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed creating account: %v", err)
		return nil, err
	}

	res := model.AccountResponse{
		Email:     acc.Email,
		UpdateAt:  acc.UpdatedAt.Time,
		CreatedAt: acc.CreatedAt.Time,
	}

	return &res, nil
}
