package service

import (
	"context"
	"testing"

	"miniWiki/internal/domain/account/model"
	"miniWiki/pkg/security"

	"github.com/stretchr/testify/suite"
)

var (
	testAccountId     = 1
	testUpdateAccount = model.UpdateAccount{
		Name: &name,
	}
	testUpdateAccountRequest = model.UpdateAccountRequest{
		AccountId: testAccountId,
		Account:   testUpdateAccount,
	}
)

type UpdateAccountTestSuite struct {
	suite.Suite
	hash    *security.HashMock
	aRepo   *accountRepositoryMock
	ctx     context.Context
	service *Account
}

func (s *UpdateAccountTestSuite) SetupSuite() {
	s.hash = &security.HashMock{}
	s.aRepo = &accountRepositoryMock{}
	s.ctx = context.Background()

	s.service = NewAccount(s.aRepo, nil, s.hash, nil)
}

func (s *UpdateAccountTestSuite) TestUpdateAccount_HashError() {
	req := model.UpdateAccountRequest{
		AccountId: 1,
		Account: model.UpdateAccount{
			Password: &password,
		},
	}
	s.hash.On("GenerateFormatted", password).
		Return("", testError).
		Once()

	err := s.service.UpdateAccount(s.ctx, req)
	s.Error(err)
}

func (s *UpdateAccountTestSuite) TestUpdateAccount_InvalidPassword() {
	pass := "pass"
	req := model.UpdateAccountRequest{
		AccountId: 1,
		Account: model.UpdateAccount{
			Password: &pass,
		},
	}
	err := s.service.UpdateAccount(s.ctx, req)
	s.Error(err)
}

func (s *UpdateAccountTestSuite) TestUpdateAccount_RepoError() {
	s.hash.On("GenerateFormatted", password).
		Return(password, nil).
		Once()
	s.aRepo.On("UpdateAccount", testAccountId, testUpdateAccount).
		Return(testError).
		Once()

	err := s.service.UpdateAccount(s.ctx, testUpdateAccountRequest)
	s.Error(err)
}

func (s *UpdateAccountTestSuite) TestUpdateAccount_Successful() {
	s.hash.On("GenerateFormatted", password).
		Return(password, nil).
		Once()
	s.aRepo.On("UpdateAccount", testAccountId, testUpdateAccount).
		Return(nil).
		Once()

	err := s.service.UpdateAccount(s.ctx, testUpdateAccountRequest)
	s.NoError(err)
}

func (s *UpdateAccountTestSuite) AfterTest(_, _ string) {
	s.aRepo.AssertExpectations(s.T())
}

func TestUpdateAccountSuite(t *testing.T) {
	suite.Run(t, new(UpdateAccountTestSuite))
}
