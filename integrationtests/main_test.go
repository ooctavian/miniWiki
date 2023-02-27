package integrationtests

import (
	"context"
	"net/http/httptest"

	"miniWiki/app"
	"miniWiki/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

// Integration tests are now on hold till I write user

type IntegrationTestSuite struct {
	suite.Suite
	ctx context.Context
	db  *pgxpool.Pool
	srv *httptest.Server
	clt *testClient
}

func (s *IntegrationTestSuite) SetupSuite() {
	err := godotenv.Load("../.env")
	s.NoError(err)
	cfg, err := config.InitConfig()
	s.NoError(err)

	s.ctx = context.Background()
	db, err := pgxpool.Connect(s.ctx, cfg.Database.DatabaseURL)
	s.NoError(err)
	s.db = db
	s.srv = httptest.NewServer(app.InitRouter(db, *cfg))
	s.clt = newTestClient(s.srv.URL, s.T())
}
