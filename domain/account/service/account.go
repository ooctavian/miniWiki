package service

import (
	"miniWiki/domain/account/query"
	rQuery "miniWiki/domain/resource/query"
	"miniWiki/security"
)

type Account struct {
	accountQuerier  query.Querier
	resourceQuerier rQuery.Querier
	hash            security.Hash
}

func NewAccount(querier query.Querier, resourceQuerier rQuery.Querier, hashAlgorithm security.Hash) *Account {
	account := &Account{
		accountQuerier:  querier,
		resourceQuerier: resourceQuerier,
		hash:            hashAlgorithm,
	}
	return account
}
