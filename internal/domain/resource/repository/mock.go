package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ResourceRepositoryMock struct {
	mock.Mock
}

func (r *ResourceRepositoryMock) CountCategoryResources(_ context.Context, id int) (int64, error) {
	args := r.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (r *ResourceRepositoryMock) MakeResourcesPrivate(_ context.Context, id int) error {
	args := r.Called(id)
	return args.Error(0)
}
