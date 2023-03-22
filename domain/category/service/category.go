package service

import (
	"gorm.io/gorm"
)

type Category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	category := &Category{
		db: db,
	}
	return category
}
