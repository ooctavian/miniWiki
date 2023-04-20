package service

import (
	"context"
	"errors"
	"testing"

	model2 "miniWiki/internal/domain/category/model"
	"miniWiki/internal/domain/resource/model"

	"github.com/stretchr/testify/suite"
)

var (
	testCreateResource = model.CreateResource{
		Title:       testTitle,
		Description: testDescription,
		Link:        testLink,
		State:       testState,
		CategoryId:  testCategoryId,
		AuthorId:    testAuthorId,
	}
	testCreateResourceRequest = model.CreateResourceRequest{
		Resource:  testCreateResource,
		AccountId: testAccountId,
	}
	testCategoryName   = "lorem"
	testCreateCategory = model2.CreateCategoryRequest{
		Category: model2.CreateCategory{
			Title: testCategoryName,
		},
	}
	testError = errors.New("test")
)

type CreateResourceSuite struct {
	suite.Suite
	ctx      context.Context
	rRepo    *resourceRepositoryMock
	cService *categoryServiceMock

	service *Resource
}

func (s *CreateResourceSuite) SetupSuite() {
	s.ctx = context.Background()
	s.rRepo = &resourceRepositoryMock{}
	s.cService = &categoryServiceMock{}

	s.service = NewResource(s.rRepo, s.cService, nil)
}

func (s *CreateResourceSuite) TestCreateResource_Succesful() {
	s.rRepo.On("InsertResource", testCreateResource).
		Return(testResourceIdUint, nil).
		Once()
	id, err := s.service.CreateResource(s.ctx, testCreateResourceRequest)
	s.NoError(err)
	s.Equal(id, &testResourceIdUint)
}

func (s *CreateResourceSuite) TestCreateResource_NewCategory_Succesful() {
	resource := model.CreateResource{
		Title:        testTitle,
		Description:  testDescription,
		Link:         testLink,
		State:        testState,
		CategoryName: &testCategoryName,
		AuthorId:     testAuthorId,
		CategoryId:   testCategoryId,
	}

	request := model.CreateResourceRequest{
		Resource:  resource,
		AccountId: testAccountId,
	}

	s.cService.On("CreateCategory", testCreateCategory).
		Return(&testCategoryId, nil).
		Once()
	s.rRepo.On("InsertResource", resource).
		Return(testResourceIdUint, nil).
		Once()
	id, err := s.service.CreateResource(s.ctx, request)
	s.NoError(err)
	s.Equal(id, &testResourceIdUint)
}

func (s *CreateResourceSuite) TestCreateResource_NewCategory_Failed() {
	resource := model.CreateResource{
		Title:        testTitle,
		Description:  testDescription,
		Link:         testLink,
		State:        testState,
		CategoryName: &testCategoryName,
		AuthorId:     testAuthorId,
		CategoryId:   testCategoryId,
	}

	request := model.CreateResourceRequest{
		Resource:  resource,
		AccountId: testAccountId,
	}

	s.cService.On("CreateCategory", testCreateCategory).
		Return(&testCategoryId, testError).
		Once()
	id, err := s.service.CreateResource(s.ctx, request)
	s.Error(err)
	s.Nil(id)
}

func (s *CreateResourceSuite) AfterTest(_, _ string) {
	s.rRepo.AssertExpectations(s.T())
	s.cService.AssertExpectations(s.T())
}

func TestCreateResourceSuite(t *testing.T) {
	suite.Run(t, new(CreateResourceSuite))
}
