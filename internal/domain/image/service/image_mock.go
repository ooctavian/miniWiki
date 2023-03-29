package service

import (
	"context"
	"io"

	"miniWiki/internal/domain/image/model"

	"github.com/stretchr/testify/mock"
)

type ImageMock struct {
	mock.Mock
}

func (m *ImageMock) Upload(_ context.Context, request model.UploadRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *ImageMock) Download(_ context.Context, request model.DownloadRequest) (io.Reader, error) {
	args := m.Called(request)
	return args.Get(0).(io.Reader), args.Error(1)
}
