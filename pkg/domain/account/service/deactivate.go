package service

import (
	"context"

	"miniWiki/pkg/domain/account/model"
	rmodel "miniWiki/pkg/domain/resource/model"

	"github.com/sirupsen/logrus"
)

func (s *Account) DeactivateAccount(ctx context.Context, request model.DeactivateAccountRequest) error {
	err := s.db.Model(&rmodel.Resource{}).
		Where("category_id = ?", request.AccountId).
		Update("state", "PRIVATE").
		Error
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed deactivating account: %v", err)
		return err
	}

	err = s.db.Model(&model.Account{}).
		Where("account_id = ?", request.AccountId).
		Update("active", false).
		Error
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed changing account status: %v", err)
		return err
	}
	return nil
}
