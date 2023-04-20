package repository

import (
	"context"

	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	mCategory = model.Category{}
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r CategoryRepository) CreateCategory(ctx context.Context, category model.CreateCategory) (int, error) {
	err := r.db.WithContext(ctx).Create(&category).Error
	if err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (r CategoryRepository) GetCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	var categories []model.Category
	err := r.db.WithContext(ctx).
		Scopes(pagination.Paginate(categories, r.db)).
		Find(&categories).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = categories
	return pagination, nil
}

func (r CategoryRepository) GetCategory(ctx context.Context, id int) (model.Category, error) {
	var c model.Category
	err := r.db.WithContext(ctx).First(&c, id).Error
	if err != nil {
		return model.Category{}, err
	}

	return c, nil
}

func (r CategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Delete(mCategory, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r CategoryRepository) CountCategories(ctx context.Context, id int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&mCategory).
		Where("parent_id = ?", id).
		Count(&count).Error
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return 0, err
	}
	return count, nil
}
