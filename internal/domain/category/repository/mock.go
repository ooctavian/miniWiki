package repository

import (
	"context"

	"miniWiki/internal/domain/category/model"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (s *CategoryRepositoryMock) CreateCategory(_ context.Context, category model.CreateCategory) (model.CreateCategory, error) {
	args := s.Called(category)
	return args.Get(0).(model.CreateCategory), args.Error(1)
}

func (s *CategoryRepositoryMock) GetCategories(_ context.Context) ([]model.Category, error) {
	args := s.Called()
	return args.Get(0).([]model.Category), args.Error(1)
}

func (s *CategoryRepositoryMock) GetCategory(_ context.Context, id int) (model.Category, error) {
	args := s.Called(id)
	return args.Get(0).(model.Category), args.Error(1)
}

func (s *CategoryRepositoryMock) DeleteCategory(_ context.Context, id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *CategoryRepositoryMock) CountCategories(_ context.Context, id int) (int64, error) {
	args := s.Called(id)
	return args.Get(0).(int64), args.Error(1)
}
