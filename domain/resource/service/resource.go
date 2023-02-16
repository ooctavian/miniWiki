package service

import (
	"miniWiki/domain/resource/query"
)

type Resource struct {
	resourceQuerier *query.DBQuerier
}

func NewResource(querier *query.DBQuerier) *Resource {
	resource := &Resource{}
	resource.resourceQuerier = querier
	return resource
}
