package service

import (
	"context"
	"io"

	cQuery "miniWiki/domain/category/query"
	model2 "miniWiki/domain/image/model"
	rQuery "miniWiki/domain/resource/query"
)

type imageService interface {
	Upload(ctx context.Context, request model2.UploadRequest) error
	Download(ctx context.Context, request model2.DownloadRequest) (io.Reader, error)
}

type Resource struct {
	resourceQuerier rQuery.Querier
	categoryQuerier cQuery.Querier
	imageService    imageService
}

func NewResource(rQuerier rQuery.Querier, cQuerier cQuery.Querier, service imageService) *Resource {
	resource := &Resource{
		resourceQuerier: rQuerier,
		categoryQuerier: cQuerier,
		imageService:    service,
	}
	return resource
}
