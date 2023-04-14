package service

import (
	"context"

	model2 "miniWiki/internal/domain/category/model"
	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"

	"github.com/stretchr/testify/mock"
)

type resourceRepositoryMock struct {
	mock.Mock
}

func (m *resourceRepositoryMock) GetResourceById(_ context.Context, id int) (*model.Resource, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Resource), args.Error(1)
}

func (m *resourceRepositoryMock) GetResources(_ context.Context, accountId int, pagination utils.Pagination, filters model.GetResourcesFilters) (utils.Pagination, error) {
	args := m.Called(accountId, pagination, filters)
	return args.Get(0).(utils.Pagination), args.Error(1)
}

func (m *resourceRepositoryMock) DeleteResourceById(_ context.Context, resourceId uint, accountId uint) error {
	args := m.Called(resourceId, accountId)
	return args.Error(0)
}

func (m *resourceRepositoryMock) InsertResource(_ context.Context, resource model.CreateResource) (uint, error) {
	args := m.Called(resource)
	return args.Get(0).(uint), args.Error(1)
}

func (m *resourceRepositoryMock) UpdateResource(_ context.Context, request model.UpdateResourceRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *resourceRepositoryMock) UpdateResourcePicture(_ context.Context, resourceId int, accountId int, path string) error {
	args := m.Called(resourceId, accountId, path)
	return args.Error(0)
}

type categoryServiceMock struct {
	mock.Mock
}

func (m *categoryServiceMock) CreateCategory(_ context.Context, request model2.CreateCategoryRequest) (*int, error) {
	args := m.Called(request)
	return args.Get(0).(*int), args.Error(1)
}
