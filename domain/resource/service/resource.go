package service

import (
	cQuery "miniWiki/domain/category/query"
	iService "miniWiki/domain/image/service"
	rQuery "miniWiki/domain/resource/query"
)

type Resource struct {
	resourceQuerier rQuery.Querier
	categoryQuerier cQuery.Querier
	imageService    iService.ImageService
}

func NewResource(rQuerier rQuery.Querier, cQuerier cQuery.Querier, service iService.ImageService) *Resource {
	resource := &Resource{
		resourceQuerier: rQuerier,
		categoryQuerier: cQuerier,
		imageService:    service,
	}
	return resource
}
