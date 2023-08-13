package service

import (
	"context"

	"miniWiki/internal/domain/filemanager/model"
)

type FileManager struct {
	Destination string
}

func NewFileManager(destination string) *FileManager {
	return &FileManager{
		Destination: destination,
	}
}

type Uploader interface {
	Upload(ctx context.Context, request model.UploadRequest) error
}
