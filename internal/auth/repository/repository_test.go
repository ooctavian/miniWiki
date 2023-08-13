package repository

import (
	"context"
	"miniWiki/internal/auth/model"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	testSessionId = "test"
	testAccountId = 1
	testIpAddress = "test"
	testUserAgent = "test"
	testExpireAt  = time.Time{}
	testSession   = model.Session{
		SessionID: testSessionId,
		AccountID: testAccountId,
		IpAddress: testIpAddress,
		UserAgent: testUserAgent,
		ExpireAt:  testExpireAt,
	}
)

type AuthRepositorySuite struct {
	suite.Suite
	ctx        context.Context
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository *AuthRepository
}

func (s *AuthRepositorySuite) SetupSuite() {
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
	s.repository = NewAuthRepository(s.DB)
}

func (s *AuthRepositorySuite) AfterTest(_, _ string) {
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *AuthRepositorySuite) Test_Repository_GetSession() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "session" WHERE session_id = $1 LIMIT 1`)).
		WithArgs(testSessionId).
		WillReturnRows(
			sqlmock.NewRows([]string{"session_id", "account_id", "ip_address", "user_agent", "expire_at"}).
				AddRow(testSessionId, testAccountId, testIpAddress, testUserAgent, testExpireAt),
		)
	session, err := s.repository.GetSession(s.ctx, testSessionId)
	s.NoError(err)
	s.Equal(&testSession, session)
}

func (s *AuthRepositorySuite) Test_Repository_DeleteSession() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "session" WHERE "session_id" = $1`)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteSession(s.ctx, testSessionId)
	s.NoError(err)
}

func (s *AuthRepositorySuite) Test_Repository_CreateSession() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "session" ("session_id","account_id","ip_address","user_agent","expire_at") VALUES ($1,$2,$3,$4,$5)`)).
		WithArgs(testSessionId, testAccountId, testIpAddress, testUserAgent, testExpireAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.CreateSession(s.ctx, testSession)
	s.NoError(err)
}

func (s *AuthRepositorySuite) Test_Repository_UpdateSession() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "session" SET "session_id"=$1,"ip_address"=$2 WHERE session_id = $3`)).
		WithArgs(testSessionId, testIpAddress, testSessionId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateSession(s.ctx, testSessionId, model.Session{
		SessionID: testSessionId,
		IpAddress: testIpAddress,
	})
	s.NoError(err)
}

func TestAuthRepositorySuite(t *testing.T) {
	suite.Run(t, new(AuthRepositorySuite))
}
