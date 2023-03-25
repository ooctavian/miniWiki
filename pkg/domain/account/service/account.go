package service

import (
	iService "miniWiki/pkg/domain/image/service"
	"miniWiki/pkg/security"

	"gorm.io/gorm"
)

type Account struct {
	hash         security.Hash
	db           *gorm.DB
	imageService iService.ImageService
}

func NewAccount(db *gorm.DB, hashAlgorithm security.Hash, imageService iService.ImageService) *Account {
	account := &Account{
		db:           db,
		hash:         hashAlgorithm,
		imageService: imageService,
	}
	return account
}
