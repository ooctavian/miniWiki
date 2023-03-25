package service

import (
	"context"
	"io"

	"miniWiki/internal/domain/image/model"
)

type Image struct {
	Destination string
}

func NewImage(destination string) *Image {
	return &Image{
		Destination: destination,
	}
}

type ImageService interface {
	Upload(ctx context.Context, request model.UploadRequest) error
	Download(ctx context.Context, request model.DownloadRequest) (io.Reader, error)
}
