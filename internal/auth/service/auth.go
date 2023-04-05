package service

import (
	"context"

	"miniWiki/internal/auth/model"
	model2 "miniWiki/internal/domain/account/model"
	"miniWiki/pkg/security"
)

type accountRepositoryInterface interface {
	GetAccount(ctx context.Context, id int) (model2.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (model2.Account, error)
	UpdateAccountStatus(ctx context.Context, id int, status bool) error
}

type Auth struct {
	accountRepository accountRepositoryInterface
	authRepository    authRepositoryInterface
	hash              security.Hash
}

type authRepositoryInterface interface {
	GetSession(ctx context.Context, sessionID string) (*model.Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
	UpdateSession(ctx context.Context, sessionID string, session model.Session) error
	CreateSession(ctx context.Context, session model.Session) error
}

func NewAuth(
	accountRepository accountRepositoryInterface,
	authRepository authRepositoryInterface,
	hash security.Hash,
) *Auth {
	return &Auth{
		accountRepository: accountRepository,
		authRepository:    authRepository,
		hash:              hash,
	}
}
