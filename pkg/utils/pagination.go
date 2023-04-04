package utils

import (
	"math"

	"gorm.io/gorm"
)

// Stolen from: https://dev.to/rafaelgfirmino/pagination-using-gorm-scopes-3k5f

// swagger:model Pagination
type Pagination struct {
	// Maximum of elements per page
	// example: 10
	Limit int `json:"limit,omitempty" schema:"limit"`
	// Number of page
	// example: 1
	Page int `json:"page,omitempty" schema:"page"`
	// Total number of elements
	// example: 200
	TotalRows int64 `json:"total_rows"`
	// Total number of pages
	// example: 10
	TotalPages int `json:"total_pages"`
	// The data returned accordingly to the parameters
	Data interface{} `json:"data"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) Paginate(value interface{}, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	p.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.Limit)))
	p.TotalPages = totalPages
	if totalPages < 0 {
		p.TotalPages = 1
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit())
	}
}
