package service

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"github.com/stretchr/testify/mock"
)

type ResourceRepositoryMock struct {
	mock.Mock
}

func (r *ResourceRepositoryMock) MakeResourcesPrivate(_ context.Context, id int) error {
	args := r.Called(id)
	return args.Error(0)
}

type accountRepositoryMock struct {
	mock.Mock
}

func (m *accountRepositoryMock) CreateAccount(_ context.Context, acc model.CreateAccount) error {
	args := m.Called(acc)
	return args.Error(0)
}

func (m *accountRepositoryMock) UpdateAccount(_ context.Context, id int, acc model.UpdateAccount) error {
	args := m.Called(id, acc)
	return args.Error(0)
}

func (m *accountRepositoryMock) GetAccount(_ context.Context, id int) (model.Account, error) {
	args := m.Called(id)
	return args.Get(0).(model.Account), args.Error(1)
}
