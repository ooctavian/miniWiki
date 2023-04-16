package service

import (
	"context"
	"errors"
	"testing"

	"miniWiki/internal/domain/account/model"
	"miniWiki/pkg/security"

	"github.com/stretchr/testify/suite"
)

var (
	email       = "example@mail.com"
	password    = "PaS$W0rD"
	name        = "Lorem Ipsum"
	testError   = errors.New("test")
	testAccount = model.CreateAccount{
		Email:    email,
		Password: password,
		Name:     name,
	}
	testCreateAccountRequest = model.CreateAccountRequest{Account: testAccount}
)

type CreateAccountTestSuite struct {
	suite.Suite
	hash    *security.HashMock
	aRepo   *accountRepositoryMock
	ctx     context.Context
	service *Account
}

func (s *CreateAccountTestSuite) SetupSuite() {
	s.hash = &security.HashMock{}
	s.aRepo = &accountRepositoryMock{}
	s.ctx = context.Background()

	s.service = NewAccount(s.aRepo, nil, s.hash, nil)
}

func (s *CreateAccountTestSuite) TestCreateAccount_Successful() {

	s.hash.On("GenerateFormatted", password).
		Return(password, nil).
		Once()
	s.aRepo.On("CreateAccount", testAccount).
		Return(nil).
		Once()

	err := s.service.CreateAccount(s.ctx, testCreateAccountRequest)
	s.NoError(err)
}

func (s *CreateAccountTestSuite) TestCreateAccount_RepoError() {
	s.hash.On("GenerateFormatted", password).
		Return(password, nil).
		Once()
	s.aRepo.On("CreateAccount", testAccount).
		Return(testError).
		Once()

	err := s.service.CreateAccount(s.ctx, testCreateAccountRequest)
	s.Error(err)
}

func (s *CreateAccountTestSuite) TestCreateAccount_HashError() {
	s.hash.On("GenerateFormatted", password).
		Return("", testError).
		Once()

	err := s.service.CreateAccount(s.ctx, testCreateAccountRequest)
	s.Error(err)
}

func (s *CreateAccountTestSuite) TestCreateAccount_InvalidPassword() {
	acc := model.CreateAccount{
		Email:    email,
		Password: "pass",
		Name:     name,
	}

	err := s.service.CreateAccount(s.ctx, model.CreateAccountRequest{Account: acc})
	s.Error(err)
}

func (s *CreateAccountTestSuite) AfterTest(_, _ string) {
	s.aRepo.AssertExpectations(s.T())
}

func TestCreateAccountSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountTestSuite))
}
