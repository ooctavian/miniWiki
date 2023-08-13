package service

import (
	"context"

	"miniWiki/internal/domain/filemanager/model"

	"github.com/stretchr/testify/mock"
)

type UploaderMock struct {
	mock.Mock
}

func (m *UploaderMock) Upload(_ context.Context, request model.UploadRequest) error {
	args := m.Called(request)
	return args.Error(0)
}
