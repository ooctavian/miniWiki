package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	testResourceId  = 1
	testTitle       = "lorem ipsum"
	testDescription = "lorem ipsum"
	testError       = errors.New("error")
	testAccountId   = 1
	testResource    = model.Resource{
		ID:          uint(testResourceId),
		Title:       testTitle,
		Description: testDescription,
	}
	testLink       = "https://example.com"
	testResources  = []model.Resource{testResource}
	testPagination = utils.Pagination{
		Limit:      5,
		Page:       1,
		TotalPages: 1,
		TotalRows:  1,
		Data:       testResources,
	}
	testFilters = model.GetResourcesFilters{
		Title: "lorem",
		Link:  "ipsum",
	}
	testState          = "PUBLIC"
	testCategoryId     = 1
	testCreateResource = model.CreateResource{
		Title:       testTitle,
		Description: testDescription,
		Link:        testLink,
		AuthorId:    testAccountId,
		State:       testState,
		CategoryId:  testCategoryId,
	}
	testUpdateResource = model.UpdateResource{
		Title: &testTitle,
	}
	testUpdateResourceRequest = model.UpdateResourceRequest{
		ResourceId: testResourceId,
		AccountId:  testAccountId,
		Resource:   testUpdateResource,
	}
	testPicturePath = "/image.png"
)

type ResourceRepositorySuite struct {
	suite.Suite
	ctx        context.Context
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository *ResourceRepository
}

func (s *ResourceRepositorySuite) SetupSuite() {
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
	s.repository = NewResourceRepository(s.DB)
}

func (s *ResourceRepositorySuite) AfterTest(_, _ string) {
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *ResourceRepositorySuite) Test_Repository_GetResourceById_Successful() {
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resource" WHERE "resource"."resource_id" = $1 LIMIT 1`)).
		WithArgs(testResourceId).
		WillReturnRows(
			sqlmock.NewRows([]string{"resource_id", "title", "description"}).
				AddRow(testResourceId, testTitle, testDescription),
		)
	res, err := s.repository.GetResourceById(s.ctx, testResourceId)
	s.NoError(err)
	s.Equal(res, &testResource)
}

func (s *ResourceRepositorySuite) Test_Repository_GetResourceById_Failed() {
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resource" WHERE "resource"."resource_id" = $1 LIMIT 1`)).
		WithArgs(testResourceId).
		WillReturnError(testError)
	res, err := s.repository.GetResourceById(s.ctx, testResourceId)
	s.Error(err)
	s.Nil(res)
}

func (s *ResourceRepositorySuite) Test_Repository_GetResources_Successful() {
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "resource"`)).
		WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resource" WHERE ($1::TEXT IS NULL OR title LIKE '%' || $2 || '%')
		AND ($3::TEXT IS NULL OR link LIKE '%'|| $4 ||'%')
		AND ($5 IS NULL or category_id = ANY($6))
		AND (state = 'PUBLIC' OR author_id = $7) LIMIT 5`)).
		WithArgs(testFilters.Title, testFilters.Title,
			testFilters.Link, testFilters.Link,
			nil, nil,
			testAccountId,
		).
		WillReturnRows(
			sqlmock.NewRows([]string{"resource_id", "title", "description"}).
				AddRow(testResourceId, testTitle, testDescription),
		)
	res, err := s.repository.GetResources(s.ctx, testAccountId, testPagination, testFilters)
	s.NoError(err)
	s.Equal(testPagination, res)
}

func (s *ResourceRepositorySuite) Test_Repository_GetResources_Failed() {
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "resource"`)).
		WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resource" WHERE ($1::TEXT IS NULL OR title LIKE '%' || $2 || '%')
		AND ($3::TEXT IS NULL OR link LIKE '%'|| $4 ||'%')
		AND ($5 IS NULL or category_id = ANY($6))
		AND (state = 'PUBLIC' OR author_id = $7) LIMIT 5`)).
		WithArgs(testFilters.Title, testFilters.Title,
			testFilters.Link, testFilters.Link,
			nil, nil,
			testAccountId,
		).
		WillReturnError(testError)
	_, err := s.repository.GetResources(s.ctx, testAccountId, testPagination, testFilters)
	s.Error(err)
}

func (s *ResourceRepositorySuite) Test_Repository_InsertResource_Successful() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "resource" ("title","description","link","category_id","state","author_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "resource_id"`)).
		WithArgs(testTitle, testDescription, testLink, testCategoryId, testState, testAccountId).
		WillReturnRows(sqlmock.NewRows([]string{"resource_id"}).AddRow(testResourceId))
	s.mock.ExpectCommit()
	id, err := s.repository.InsertResource(s.ctx, testCreateResource)
	s.NoError(err)
	s.Equal(uint(testResourceId), id)
}

func (s *ResourceRepositorySuite) Test_Repository_InsertResource_Failed() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "resource" ("title","description","link","category_id","state","author_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "resource_id"`)).
		WillReturnError(testError)
	_, err := s.repository.InsertResource(s.ctx, testCreateResource)
	s.Error(err)
}

func (s *ResourceRepositorySuite) Test_Repository_UpdateResource_Successful() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "resource" SET "title"=$1 WHERE resource_id = $2 AND author_id = $3`)).
		WithArgs(testTitle, testResourceId, testAccountId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateResource(s.ctx, testUpdateResourceRequest)
	s.NoError(err)
}

func (s *ResourceRepositorySuite) Test_Repository_UpdateResource_Failed() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "resource" SET "title"=$1 WHERE resource_id = $2 AND author_id = $3`)).
		WithArgs(testTitle, testResourceId, testAccountId).
		WillReturnError(testError)
	err := s.repository.UpdateResource(s.ctx, testUpdateResourceRequest)
	s.Error(err)
}

func (s *ResourceRepositorySuite) Test_Repository_DeleteResource_Successful() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resource" WHERE resource_id = $1 AND author_id = $2`)).
		WithArgs(testResourceId, testAccountId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteResourceById(s.ctx, uint(testResourceId), uint(testAccountId))
	s.NoError(err)
}

func (s *ResourceRepositorySuite) Test_Repository_DeleteResource_Failed() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resource" WHERE resource_id = $1 AND author_id = $2`)).
		WithArgs(testResourceId, testAccountId).
		WillReturnError(testError)
	err := s.repository.DeleteResourceById(s.ctx, uint(testResourceId), uint(testAccountId))
	s.Error(err)
}

func (s *ResourceRepositorySuite) Test_Repository_UpdateResourcePicture_Successful() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "resource" SET "picture_url"=$1,"updated_at"=$2 WHERE resource_id = $3 AND author_id = $4`)).
		WithArgs(testPicturePath, utils.AnyTime{}, testResourceId, testAccountId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateResourcePicture(s.ctx, testResourceId, testAccountId, testPicturePath)
	s.NoError(err)
}

func (s *ResourceRepositorySuite) Test_Repository_CountCategoryResources_Successful() {
	testCount := 10
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "resource" WHERE category_id = $1`)).
		WithArgs(testCategoryId).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(testCount))
	count, err := s.repository.CountCategoryResources(s.ctx, testResourceId)
	s.NoError(err)
	s.Equal(int64(testCount), count)
}

func (s *ResourceRepositorySuite) Test_Repository_MakeResourcesPrivate_Successful() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "resource" SET "state"=$1,"updated_at"=$2 WHERE author_id = $3`)).
		WithArgs("PRIVATE", utils.AnyTime{}, testAccountId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.MakeResourcesPrivate(s.ctx, testAccountId)
	s.NoError(err)
}

func TestResourceRepositorySuite(t *testing.T) {
	suite.Run(t, new(ResourceRepositorySuite))
}
