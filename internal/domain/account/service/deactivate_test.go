package service

import (
	"context"
	"testing"

	"miniWiki/internal/domain/account/model"

	"github.com/stretchr/testify/suite"
)

var (
	testDeactivateAccountRequest = model.DeactivateAccountRequest{
		AccountId: testAccountId,
	}
	testInactive = false
)

type DeactivateAccountTestSuite struct {
	suite.Suite
	aRepo   *accountRepositoryMock
	rRepo   *ResourceRepositoryMock
	ctx     context.Context
	service *Account
}

func (s *DeactivateAccountTestSuite) SetupSuite() {
	s.aRepo = &accountRepositoryMock{}
	s.rRepo = &ResourceRepositoryMock{}
	s.ctx = context.Background()

	s.service = NewAccount(s.aRepo, s.rRepo, nil, nil)
}

func (s *DeactivateAccountTestSuite) TestDeactivateAccount_Successful() {
	s.rRepo.On("MakeResourcesPrivate", testAccountId).
		Return(nil).
		Once()
	s.aRepo.On("UpdateAccount", testAccountId, model.UpdateAccount{Active: &testInactive}).
		Return(nil).
		Once()

	err := s.service.DeactivateAccount(s.ctx, testDeactivateAccountRequest)
	s.NoError(err)
}

func (s *DeactivateAccountTestSuite) TestDeactivateAccount_ResourceRepoFail() {
	s.rRepo.On("MakeResourcesPrivate", testAccountId).
		Return(testError).
		Once()

	err := s.service.DeactivateAccount(s.ctx, testDeactivateAccountRequest)
	s.Error(err)
}

func (s *DeactivateAccountTestSuite) TestDeactivateAccount_AccountRepoFail() {
	s.rRepo.On("MakeResourcesPrivate", testAccountId).
		Return(nil).
		Once()
	s.aRepo.On("UpdateAccount", testAccountId, model.UpdateAccount{Active: &testInactive}).
		Return(testError).
		Once()

	err := s.service.DeactivateAccount(s.ctx, testDeactivateAccountRequest)
	s.Error(err)
}

func (s *DeactivateAccountTestSuite) AfterTest(_, _ string) {
	s.aRepo.AssertExpectations(s.T())
	s.rRepo.AssertExpectations(s.T())
}

func TestDeactivateAccountSuite(t *testing.T) {
	suite.Run(t, new(DeactivateAccountTestSuite))
}
