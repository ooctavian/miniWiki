package service

import (
	aQuery "miniWiki/pkg/domain/account/query"
	"miniWiki/pkg/domain/auth/query"
	"miniWiki/pkg/security"
)

type Auth struct {
	sessionQuerier query.Querier
	accountQuerier aQuery.Querier
	hash           security.Hash
}

func NewAuth(
	sessionQuerier query.Querier,
	accountQuerier aQuery.Querier,
	hash security.Hash,
) *Auth {
	return &Auth{
		accountQuerier: accountQuerier,
		sessionQuerier: sessionQuerier,
		hash:           hash,
	}
}
