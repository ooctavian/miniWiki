package controller

import (
	"context"

	"miniWiki/domain/resource/model"

	"github.com/go-chi/chi/v5"
)

type resourceService interface {
	CreateResource(ctx context.Context, request model.CreateResourceRequest) error
	DeleteResource(ctx context.Context, request model.DeleteResourceRequest) error
	GetResources(ctx context.Context, request model.GetResourcesRequest) ([]model.ResourceResponse, error)
	GetResource(ctx context.Context, request model.GetResourceRequest) (*model.ResourceResponse, error)
	UpdateResource(ctx context.Context, request model.UpdateResourceRequest) error
}

func MakeResourceRouter(r chi.Router, service resourceService) {
	r.Post("/", createResourceHandler(service))
	r.Patch("/{id}", updateResourceHandler(service))
	r.Delete("/{id}", deleteResourceHandler(service))
	r.Get("/{id}", getResourceHandler(service))
	r.Get("/", getResourcesHandler(service))
}
