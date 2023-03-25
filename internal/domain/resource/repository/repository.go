package repository

import (
	"context"

	"miniWiki/internal/domain/resource/model"

	"gorm.io/gorm"
)

var (
	resource = model.Resource{}
)

type ResourceRepository struct {
	db *gorm.DB
}

type ResourceRepositoryInterface interface {
	CountCategoryResources(ctx context.Context, id int) (int64, error)
	MakeResourcesPrivate(ctx context.Context, id int) error
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{
		db: db,
	}
}

func (r ResourceRepository) CountCategoryResources(ctx context.Context, id int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(resource).
		Where("category_id = ?", id).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r ResourceRepository) MakeResourcesPrivate(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).
		Model(resource).
		Where("account_id = ?", id).
		Update("state", "PRIVATE").
		Error
	return err
}
