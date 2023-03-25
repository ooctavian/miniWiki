package repository

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"gorm.io/gorm"
)

var (
	account = model.Account{}
)

type AccountRepository struct {
	db *gorm.DB
}

type AccountRepositoryInterface interface {
	CreateAccount(ctx context.Context, acc model.CreateAccount) error
	UpdateAccount(ctx context.Context, id int, acc model.UpdateAccount) error
	GetAccount(ctx context.Context, id int) (model.Account, error)
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r AccountRepository) CreateAccount(ctx context.Context, acc model.CreateAccount) error {
	err := r.db.WithContext(ctx).Create(&acc).Error
	return err
}

func (r AccountRepository) UpdateAccount(ctx context.Context, id int, acc model.UpdateAccount) error {
	err := r.db.WithContext(ctx).
		Model(&model.Account{}).
		Where("account_id = ?", id).
		Updates(acc).
		Error
	return err
}

func (r AccountRepository) GetAccount(ctx context.Context, id int) (model.Account, error) {
	acc := model.Account{}
	err := r.db.WithContext(ctx).Model(&model.Account{}).First(&account, id).Error
	return acc, err
}
