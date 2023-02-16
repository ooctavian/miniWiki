package service

import (
	"miniWiki/domain/category/query"
)

type Category struct {
	categoryQuerier *query.DBQuerier
}

func NewCategory(querier *query.DBQuerier) *Category {
	category := &Category{}
	category.categoryQuerier = querier
	return category
}
