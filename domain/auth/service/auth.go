package service

import (
	"miniWiki/security"

	aQuery "miniWiki/domain/account/query"
	"miniWiki/domain/auth/query"
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
