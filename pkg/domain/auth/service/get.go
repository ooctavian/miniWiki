package service

import (
	"context"

	"miniWiki/pkg/domain/auth/query"
)

func (s *Auth) GetSession(ctx context.Context, sessionId string) (query.GetSessionRow, error) {
	return s.sessionQuerier.GetSession(ctx, sessionId)
}

func (s *Auth) GetAccountStatus(ctx context.Context, accountId int) (*bool, error) {
	return s.accountQuerier.GetAccountStatus(ctx, accountId)
}
