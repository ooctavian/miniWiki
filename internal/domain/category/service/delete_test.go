package service_test

import (
	"context"
	"testing"

	"miniWiki/internal/domain/category/model"
	"miniWiki/internal/domain/category/service"
	"miniWiki/pkg/transport"

	"github.com/stretchr/testify/suite"
)

var (
	testDeleteCategoryRequest = model.DeleteCategoryRequest{
		CategoryId: categoryId,
	}
)

type DeleteCategorySuite struct {
	suite.Suite
	ctx     context.Context
	cRepo   *CategoryRepositoryMock
	rRepo   *ResourceRepositoryMock
	service *service.Category
}

func (s *DeleteCategorySuite) SetupSuite() {
	s.cRepo = &CategoryRepositoryMock{}
	s.rRepo = &ResourceRepositoryMock{}
	s.ctx = context.Background()

	s.service = service.NewCategory(s.cRepo, s.rRepo)
}

func (s *DeleteCategorySuite) TestDeleteCategory_Successful() {
	s.rRepo.On("CountCategoryResources", categoryId).
		Return(int64(0), nil).
		Once()
	s.cRepo.On("CountCategories", categoryId).
		Return(int64(0), nil).
		Once()
	s.cRepo.On("DeleteCategory", categoryId).
		Return(nil).
		Once()
	err := s.service.DeleteCategory(s.ctx, testDeleteCategoryRequest)
	s.NoError(err)
}

func (s *DeleteCategorySuite) TestDeleteCategory_CountCategoryResources_Error() {
	s.rRepo.On("CountCategoryResources", categoryId).
		Return(int64(0), testError).
		Once()

	err := s.service.DeleteCategory(s.ctx, testDeleteCategoryRequest)
	s.Error(err)
}

func (s *DeleteCategorySuite) TestCreateCategory_CategoryHasResources() {
	s.rRepo.On("CountCategoryResources", categoryId).
		Return(int64(1), nil).
		Once()

	err := s.service.DeleteCategory(s.ctx, testDeleteCategoryRequest)
	s.Equal(transport.ForbiddenError{}, err)
}

func (s *DeleteCategorySuite) TestDeleteCategory_CountCategories_Failed() {
	s.rRepo.On("CountCategoryResources", categoryId).
		Return(int64(0), nil).
		Once()
	s.cRepo.On("CountCategories", categoryId).
		Return(int64(0), testError).
		Once()

	err := s.service.DeleteCategory(s.ctx, testDeleteCategoryRequest)
	s.Error(err)
}

func (s *DeleteCategorySuite) TestDeleteCategory_ParentCategory() {
	s.rRepo.On("CountCategoryResources", categoryId).
		Return(int64(0), nil).
		Once()
	s.cRepo.On("CountCategories", categoryId).
		Return(int64(1), nil).
		Once()

	err := s.service.DeleteCategory(s.ctx, testDeleteCategoryRequest)
	s.Equal(transport.ForbiddenError{}, err)
}

func (s *DeleteCategorySuite) TestDeleteCategory_DeleteCategory_Fails() {
	s.rRepo.On("CountCategoryResources", categoryId).
		Return(int64(0), nil).
		Once()
	s.cRepo.On("CountCategories", categoryId).
		Return(int64(0), nil).
		Once()
	s.cRepo.On("DeleteCategory", categoryId).
		Return(testError).
		Once()
	err := s.service.DeleteCategory(s.ctx, testDeleteCategoryRequest)
	s.Error(err)
}

func TestDeleteCategorySuite(t *testing.T) {
	suite.Run(t, new(DeleteCategorySuite))
}
