package service

import (
	aRepository "miniWiki/internal/domain/account/repository"
	iService "miniWiki/internal/domain/image/service"
	rRepository "miniWiki/internal/domain/resource/repository"
	"miniWiki/pkg/security"
)

type Account struct {
	hash               security.Hash
	accountRepository  aRepository.AccountRepositoryInterface
	resourceRepository rRepository.ResourceRepositoryInterface
	imageService       iService.ImageService
}

func NewAccount(accountRepository aRepository.AccountRepositoryInterface,
	resourceRepository rRepository.ResourceRepositoryInterface,
	hashAlgorithm security.Hash,
	imageService iService.ImageService) *Account {
	account := &Account{
		accountRepository:  accountRepository,
		resourceRepository: resourceRepository,
		hash:               hashAlgorithm,
		imageService:       imageService,
	}
	return account
}
