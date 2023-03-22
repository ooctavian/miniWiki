package integrationtests_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"miniWiki/app"
	"miniWiki/config"
	model2 "miniWiki/domain/account/model"
	"miniWiki/domain/auth/model"

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
	err := godotenv.Load("../.env.test")
	s.NoError(err)
	cfg, err := config.InitConfig()
	s.NoError(err)

	s.ctx = context.Background()
	db, err := pgxpool.Connect(s.ctx, cfg.Database.DatabaseURL)
	s.NoError(err)
	s.db = db
	s.srv = httptest.NewServer(app.InitRouter(db, nil, *cfg))
	s.clt = newTestClient(s.srv.URL, s.T())
	s.CreateAccount()
}

func (s *IntegrationTestSuite) CreateAccount() {
	account := model2.CreateAccount{
		Email:    testAccountEmail,
		Password: testAccountPassword,
	}
	s.clt.Post("/account", account)
}

func (s *IntegrationTestSuite) GetAuthenticatedClient() testClient {
	credentials := model.LoginAccount{
		Email:    testAccountEmail,
		Password: testAccountPassword,
	}
	res := s.clt.Post("/login", credentials)
	clt, err := s.clt.WithCookies(res.Cookies())
	s.NoError(err)
	return clt
}

func (s *IntegrationTestSuite) TearDownTest() {
	_, err := s.db.Exec(s.ctx, "DELETE FROM resource")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "DELETE FROM category")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "DELETE FROM profile")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "ALTER SEQUENCE category_category_id_seq RESTART WITH 1;")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "UPDATE category SET category_id=nextval('category_category_id_seq');")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "ALTER SEQUENCE resource_resource_id_seq RESTART WITH 1;")
	s.NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	_, err := s.db.Exec(s.ctx, "DELETE FROM session")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "DELETE FROM account")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "UPDATE resource SET resource_id=nextval('resource_resource_id_seq');")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "ALTER SEQUENCE account_account_id_seq RESTART WITH 1;")
	s.NoError(err)
	_, err = s.db.Exec(s.ctx, "UPDATE account SET account_id=nextval('account_account_id_seq');")
	s.NoError(err)
}

func (s *IntegrationTestSuite) parseId(res *http.Response, id_position int) int {
	paths := strings.Split(res.Header.Get("Location"), "/")
	id, err := strconv.Atoi(paths[id_position])
	s.NoError(err)
	return id
}

func (s *IntegrationTestSuite) encode(v interface{}) string {
	body, err := json.Marshal(v)
	s.NoError(err)
	return string(body)
}
