package integrationtests_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"miniWiki/internal/app"
	"miniWiki/internal/auth/model"
	"miniWiki/internal/config"
	"miniWiki/pkg/security"

	"github.com/joho/godotenv"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type IntegrationTestSuite struct {
	suite.Suite
	ctx  context.Context
	db   *gorm.DB
	srv  *httptest.Server
	clt  *testClient
	hash security.Hash
}

func (s *IntegrationTestSuite) SetupSuite() {
	err := godotenv.Load("../../.env.test")
	s.NoError(err)
	cfg, err := config.InitConfig()
	s.NoError(err)
	s.hash = security.NewArgon2id(
		cfg.Argon2id.Memory,
		cfg.Argon2id.Iterations,
		cfg.Argon2id.Parallelism,
		cfg.Argon2id.SaltLength,
		cfg.Argon2id.KeyLength,
		security.GenerateRandomBytes,
	)

	s.ctx = context.Background()
	db, err := gorm.Open(postgres.Open(cfg.Database.DatabaseURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gorm_logrus.New(),
	})
	s.NoError(err)
	s.db = db
	s.srv = httptest.NewServer(app.InitRouter(db, *cfg))
	s.clt = newTestClient(s.srv.URL, s.T(), s.ctx)
	s.CreateAccount()
}

func (s *IntegrationTestSuite) CreateAccount() {
	pass, err := s.hash.GenerateFormatted(testAccountPassword)
	s.NoError(err)
	err = s.db.Exec("INSERT INTO account(email,password,name) VALUES (?,?,?);",
		testAccountEmail,
		pass,
		"Test",
	).Error
	s.NoError(err)
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
	err := s.db.Exec("DELETE FROM resource").Error
	s.NoError(err)
	err = s.db.Exec("DELETE FROM category").Error
	s.NoError(err)
	err = s.db.Exec("ALTER SEQUENCE category_category_id_seq RESTART WITH 1;").Error
	s.NoError(err)
	err = s.db.Exec("UPDATE category SET category_id=nextval('category_category_id_seq');").Error
	s.NoError(err)
	err = s.db.Exec("ALTER SEQUENCE resource_resource_id_seq RESTART WITH 1;").Error
	s.NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	err := s.db.Exec("DELETE FROM session").Error
	s.NoError(err)
	err = s.db.Exec("DELETE FROM account").Error
	s.NoError(err)
	err = s.db.Exec("UPDATE resource SET resource_id=nextval('resource_resource_id_seq');").Error
	s.NoError(err)
	err = s.db.Exec("ALTER SEQUENCE account_account_id_seq RESTART WITH 1;").Error
	s.NoError(err)
	err = s.db.Exec("UPDATE account SET account_id=nextval('account_account_id_seq');").Error
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
