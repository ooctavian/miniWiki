package repository

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
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
		Where("account_id = ?", id).
		Updates(&acc).
		Error
	return err
}

func (r AccountRepository) GetAccount(ctx context.Context, id int) (model.Account, error) {
	var acc model.Account
	err := r.db.WithContext(ctx).Model(&model.Account{}).Take(&acc, id).Error
	return acc, err
}

func (r AccountRepository) GetAccountByEmail(ctx context.Context, email string) (model.Account, error) {
	var acc model.Account
	err := r.db.WithContext(ctx).Take(&acc, "email = ?", email).Error
	return acc, err
}

func (r AccountRepository) UpdateAccountStatus(ctx context.Context, id int, status bool) error {
	err := r.db.WithContext(ctx).
		Model(&model.Account{}).
		Where("account_id = ?", id).
		Update("active", status).
		Error
	return err
}
