package controller

import (
	"context"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type resourceService interface {
	CreateResource(ctx context.Context, request model.CreateResourceRequest) (*uint, error)
	DeleteResource(ctx context.Context, request model.DeleteResourceRequest) error
	GetResources(ctx context.Context, request model.GetResourcesRequest) (utils.Pagination, error)
	GetResource(ctx context.Context, request model.GetResourceRequest) (*model.Resource, error)
	UpdateResource(ctx context.Context, request model.UpdateResourceRequest) error
	UploadResourceImage(ctx context.Context, request model.UploadResourceImageRequest) error
}

func MakeResourceRouter(r chi.Router, rService resourceService) {
	r.Get("/{id}", getResourceHandler(rService))
	r.Get("/", getResourcesHandler(rService))
	r.Post("/", createResourceHandler(rService))
	r.Patch("/{id}", updateResourceHandler(rService))
	r.Delete("/{id}", deleteResourceHandler(rService))
	r.Post("/{id}/image", uploadResourceImageHandler(rService))
}
