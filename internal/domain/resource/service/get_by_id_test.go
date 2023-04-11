package service

import (
	"context"
	"testing"

	"miniWiki/internal/domain/resource/model"

	"github.com/stretchr/testify/suite"
)

var (
	testResourceId         = 1
	testResourceIdUint     = uint(testResourceId)
	testAccountId          = 1
	testGetResourceRequest = model.GetResourceRequest{
		ResourceId: testResourceId,
		AccountId:  testAccountId,
	}
	testTitle       = "Lorem Ipsum"
	testDescription = "Lorem Ipsum"
	testLink        = "https://www.example.com/"
	testState       = "PUBLIC"
	testPictureUrl  = "example.png"
	testCategoryId  = 1
	testAuthorId    = 1
	testResource    = model.Resource{
		ID:          uint(testResourceId),
		Title:       testTitle,
		Description: testDescription,
		Link:        testLink,
		State:       testState,
		CategoryId:  &testCategoryId,
		PictureUrl:  testPictureUrl,
		AuthorId:    uint(testAuthorId),
	}
)

type GetResourceSuite struct {
	suite.Suite
	ctx   context.Context
	rRepo *resourceRepositoryMock

	service *Resource
}

func (s *GetResourceSuite) SetupSuite() {
	s.ctx = context.Background()
	s.rRepo = &resourceRepositoryMock{}

	s.service = NewResource(s.rRepo, nil, nil)
}

func (s *GetResourceSuite) TestGetResource_Successful() {
	s.rRepo.On("GetResourceById", testResourceId).
		Return(&testResource, nil).
		Once()
	res, err := s.service.GetResource(s.ctx, testGetResourceRequest)
	s.NoError(err)
	s.Equal(res, &testResource)
}

func (s *GetResourceSuite) TestGetResource_PrivateResource() {
	resource := testResource
	resource.AuthorId = 2
	resource.State = "PRIVATE"
	s.rRepo.On("GetResourceById", testResourceId).
		Return(&resource, nil).
		Once()
	res, err := s.service.GetResource(s.ctx, testGetResourceRequest)
	s.Error(err)
	s.Nil(res)
}

func TestGetResourceSuite(t *testing.T) {
	suite.Run(t, new(GetResourceSuite))
}
