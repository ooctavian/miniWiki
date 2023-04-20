package model

import (
	"time"
)

// CreateCategory Category creation request
// swagger:model CreateCategory
type CreateCategory struct {
	ID int `gorm:"column:category_id"`
	// Title of the category
	// example: backend
	Title string `json:"title" validate:"required"`
	// ID of parent category, making it a subcategory
	// example: 1
	ParentId *int `json:"parentId"`
}

func (CreateCategory) TableName() string {
	return "category"
}

type CreateCategoryRequest struct {
	Category CreateCategory
}

// UpdateCategory Category update request
// swagger:model UpdateCategory
type UpdateCategory struct {
	// Title of the category
	// example: backend
	Title string `json:"title"`
	// ID of parent category, making it a subcategory
	// example: 1
	ParentId int `json:"parentId"`
}

type UpdateCategoryRequest struct {
	CategoryId int
	Category   UpdateCategory
}

type GetCategoryRequest struct {
	CategoryId int
}

type DeleteCategoryRequest struct {
	CategoryId int
}

// swagger:model Category
type Category struct {
	// CategoryId of category
	// example: 2
	ID uint `gorm:"column:category_id" json:"categoryId"`
	// Title of the category
	// example: backend
	Title string `json:"title"`
	// ParentId of parent category, making it a subcategory
	// example: 1
	ParentId  *uint     `json:"parentId,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
