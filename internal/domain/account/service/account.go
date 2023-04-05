package service

import (
	"context"

	"miniWiki/internal/domain/account/model"
	iService "miniWiki/internal/domain/image/service"
	"miniWiki/pkg/security"
)

type resourceRepositoryInterface interface {
	MakeResourcesPrivate(ctx context.Context, id int) error
}

type accountRepositoryInterface interface {
	CreateAccount(ctx context.Context, acc model.CreateAccount) error
	UpdateAccount(ctx context.Context, id int, acc model.UpdateAccount) error
	GetAccount(ctx context.Context, id int) (model.Account, error)
}

type Account struct {
	hash               security.Hash
	accountRepository  accountRepositoryInterface
	resourceRepository resourceRepositoryInterface
	imageService       iService.ImageService
}

func NewAccount(
	accountRepository accountRepositoryInterface,
	resourceRepository resourceRepositoryInterface,
	hashAlgorithm security.Hash,
	imageService iService.ImageService,
) *Account {
	account := &Account{
		accountRepository:  accountRepository,
		resourceRepository: resourceRepository,
		hash:               hashAlgorithm,
		imageService:       imageService,
	}
	return account
}
