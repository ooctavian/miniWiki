package service_test

import (
	"context"
	"errors"
	"testing"

	"miniWiki/internal/domain/category/model"
	cRepository "miniWiki/internal/domain/category/repository"
	"miniWiki/internal/domain/category/service"

	"github.com/stretchr/testify/suite"
)

type CreateCategorySuite struct {
	suite.Suite
	ctx     context.Context
	cRepo   *cRepository.CategoryRepositoryMock
	service *service.Category
}

var (
	title      = "backend"
	categoryId = 1
)

func (s *CreateCategorySuite) SetupSuite() {
	s.cRepo = &cRepository.CategoryRepositoryMock{}
	s.ctx = context.Background()

	s.service = service.NewCategory(s.cRepo, nil)
}

func (s *CreateCategorySuite) TestCreateCategory_Successful() {
	cat := model.CreateCategory{Title: title}
	s.cRepo.On("CreateCategory", cat).
		Return(model.CreateCategory{ID: categoryId, Title: title}, nil).
		Once()
	req := model.CreateCategoryRequest{
		Category: cat,
	}
	id, err := s.service.CreateCategory(s.ctx, req)
	s.NoError(err)
	s.Equal(categoryId, *id)
}

func (s *CreateCategorySuite) TestCreateCategory_Invalid() {
	cat := model.CreateCategory{Title: title}
	s.cRepo.On("CreateCategory", cat).
		Return(model.CreateCategory{}, errors.New("test")).
		Once()
	req := model.CreateCategoryRequest{
		Category: cat,
	}

	_, err := s.service.CreateCategory(s.ctx, req)
	s.Error(err)
}

func TestCreateCategorySuite(t *testing.T) {
	suite.Run(t, new(CreateCategorySuite))
}
