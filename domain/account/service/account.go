package service

import (
	"miniWiki/domain/account/query"
	"miniWiki/security"
)

type Account struct {
	accountQuerier query.Querier
	hash           security.Hash
}

func NewAccount(querier query.Querier, hashAlgorithm security.Hash) *Account {
	account := &Account{}
	account.accountQuerier = querier
	account.hash = hashAlgorithm
	return account
}
