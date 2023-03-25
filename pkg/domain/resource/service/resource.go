package service

import (
	iService "miniWiki/pkg/domain/image/service"
	rQuery "miniWiki/pkg/domain/resource/query"
)

type Resource struct {
	resourceQuerier rQuery.Querier
	imageService    iService.ImageService
}

func NewResource(rQuerier rQuery.Querier, service iService.ImageService) *Resource {
	resource := &Resource{
		resourceQuerier: rQuerier,
		imageService:    service,
	}
	return resource
}
