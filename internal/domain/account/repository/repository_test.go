package repository

import (
	"context"
	"regexp"
	"testing"

	"miniWiki/internal/domain/account/model"
	"miniWiki/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	testEmail    = "test@example.com"
	testPassword = "pass"
	testName     = "Test"
	testId       = 1
	testAccount  = model.Account{
		ID:       testId,
		Email:    testEmail,
		Password: testPassword,
		Name:     testName,
	}
)

type AccountRepositorySuite struct {
	suite.Suite
	ctx        context.Context
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository *AccountRepository
}

func (s *AccountRepositorySuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	s.ctx = context.Background()
	s.mock = mock
	s.NoError(err)

	dialector := postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	s.NoError(err)
	s.repository = NewAccountRepository(s.DB)
}

func (s *AccountRepositorySuite) AfterTest(_, _ string) {
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *AccountRepositorySuite) Test_Repository_CreateAccount() {
	s.mock.ExpectBegin()
	s.mock.
		ExpectExec(regexp.QuoteMeta(`INSERT INTO "account" ("email","password","name","alias","picture_url") VALUES ($1,$2,$3,$4,$5)`)).
		WithArgs(testEmail, testPassword, testName, "", "").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.CreateAccount(s.ctx, model.CreateAccount{
		Email:    testEmail,
		Password: testPassword,
		Name:     testName,
	})

	s.NoError(err)
}

func (s *AccountRepositorySuite) Test_Repository_UpdateAccount() {
	s.mock.ExpectBegin()
	s.mock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "account" SET "password"=$1 WHERE account_id = $2`)).
		WithArgs(testPassword, testId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateAccount(s.ctx, testId, model.UpdateAccount{
		Password: &testPassword,
	})
	s.NoError(err)
}

func (s *AccountRepositorySuite) Test_Repository_GetAccount() {
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "account" WHERE "account"."account_id" = $1 LIMIT 1`)).
		WithArgs(testId).
		WillReturnRows(
			sqlmock.NewRows([]string{"account_id", "email", "password", "name"}).
				AddRow(testId, testEmail, testPassword, testName))

	acc, err := s.repository.GetAccount(s.ctx, testId)
	s.NoError(err)
	s.Equal(testAccount, acc)
}

func (s *AccountRepositorySuite) Test_Repository_GetAccountByEmail() {
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "account" WHERE email = $1 LIMIT 1`)).
		WithArgs(testEmail).
		WillReturnRows(
			sqlmock.NewRows([]string{"account_id", "email", "password", "name"}).
				AddRow(testId, testEmail, testPassword, testName))

	acc, err := s.repository.GetAccountByEmail(s.ctx, testEmail)
	s.NoError(err)
	s.Equal(testAccount, acc)
}

func (s *AccountRepositorySuite) Test_Repository_UpdateAccountStatus() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "account" SET "active"=$1,"updated_at"=$2 WHERE account_id = $3`)).
		WithArgs(false, utils.AnyTime{}, testId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateAccountStatus(s.ctx, testId, false)
	s.NoError(err)
}

func TestAccountRepositorySuite(t *testing.T) {
	suite.Run(t, new(AccountRepositorySuite))
}
