package repository

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m AccountRepositoryMock) CreateAccount(_ context.Context, acc model.CreateAccount) error {
	args := m.Called(acc)
	return args.Error(0)
}

func (m AccountRepositoryMock) UpdateAccount(_ context.Context, id int, acc model.UpdateAccount) error {
	args := m.Called(id, acc)
	return args.Error(0)
}

func (m AccountRepositoryMock) GetAccount(_ context.Context, id int) (model.Account, error) {
	args := m.Called(id)
	return args.Get(0).(model.Account), args.Error(1)
}
