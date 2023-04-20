package service

import (
	"context"

	"miniWiki/internal/auth/model"
)

func (s *Auth) GetSession(ctx context.Context, sessionId string) (*model.Session, error) {
	return s.authRepository.GetSession(ctx, sessionId)
}

func (s *Auth) GetAccountStatus(ctx context.Context, accountId int) (*bool, error) {
	acc, err := s.accountRepository.GetAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}
	return &acc.Active, nil
}
