package service

import (
	"context"

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
