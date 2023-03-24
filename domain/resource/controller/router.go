package controller

import (
	"context"
	"io"

	"miniWiki/domain/resource/model"

	"github.com/go-chi/chi/v5"
)

type resourceService interface {
	CreateResource(ctx context.Context, request model.CreateResourceRequest) (*model.ResourceResponse, error)
	DeleteResource(ctx context.Context, request model.DeleteResourceRequest) error
	GetResources(ctx context.Context, request model.GetResourcesRequest) ([]model.ResourceResponse, error)
	GetResource(ctx context.Context, request model.GetResourceRequest) (*model.ResourceResponse, error)
	UpdateResource(ctx context.Context, request model.UpdateResourceRequest) (*model.ResourceResponse, error)
	UploadResourceImage(ctx context.Context, request model.UploadResourceImageRequest) error
	DownloadResourceImage(ctx context.Context, request model.DownloadResourceImageRequest) (io.Reader, error)
}

func MakeResourceRouter(r chi.Router, rService resourceService) {
	r.Get("/{id}", getResourceHandler(rService))
	r.Get("/", getResourcesHandler(rService))
	r.Post("/", createResourceHandler(rService))
	r.Patch("/{id}", updateResourceHandler(rService))
	r.Delete("/{id}", deleteResourceHandler(rService))
	r.Post("/{id}/image", uploadResourceImageHandler(rService))
	r.Get("/{id}/image", downloadResourceImageHandler(rService))
}
