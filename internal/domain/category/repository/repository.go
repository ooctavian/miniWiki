package repository

import (
	"context"

	"miniWiki/internal/domain/category/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	category = model.Category{}
)

type CategoryRepository struct {
	db *gorm.DB
}

type CategoryRepositoryInterface interface {
	CreateCategory(ctx context.Context, category model.CreateCategory) (model.CreateCategory, error)
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategory(ctx context.Context, id int) (model.Category, error)
	DeleteCategory(ctx context.Context, id int) error
	CountCategories(ctx context.Context, id int) (int64, error)
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r CategoryRepository) CreateCategory(ctx context.Context, category model.CreateCategory) (model.CreateCategory, error) {
	err := r.db.WithContext(ctx).Create(category).Error
	if err != nil {
		return model.CreateCategory{}, err
	}
	return category, nil
}

func (r CategoryRepository) GetCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
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
	err := r.db.WithContext(ctx).Delete(category, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r CategoryRepository) CountCategories(ctx context.Context, id int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&category).
		Where("parent_id = ?", id).
		Count(&count).Error
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed deleting category: %v", err)
		return 0, err
	}
	return count, nil
}
