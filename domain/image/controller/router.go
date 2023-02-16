package controller

import (
	"context"
	"io"

	"miniWiki/domain/image/model"

	"github.com/go-chi/chi/v5"
)

type imageService interface {
	Upload(ctx context.Context, request model.UploadRequest) error
	Download(ctx context.Context, request model.DownloadRequest) (io.Reader, error)
}

func MakeResourceImageHandler(r chi.Router, service imageService) {
	r.Post("/{id}/image", uploadResourceImageHandler(service))
	r.Get("/{id}/image", downloadResourceImageHandler(service))
}
