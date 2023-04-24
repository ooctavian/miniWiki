package repository

import (
	"context"
	"miniWiki/internal/domain/category/model"
	"miniWiki/pkg/utils"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	testCategoryName = "lorem"
	testCategoryId   = 1
	testCategory     = model.Category{
		ID:    uint(testCategoryId),
		Title: testCategoryName,
	}
)

type CategoryRepositorySuite struct {
	suite.Suite
	ctx        context.Context
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository *CategoryRepository
}

func (s *CategoryRepositorySuite) SetupSuite() {
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
	s.repository = NewCategoryRepository(s.DB)
}

func (s *CategoryRepositorySuite) AfterTest(_, _ string) {
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *CategoryRepositorySuite) Test_Repository_CreateCategory() {
	s.mock.ExpectBegin()
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "category" ("title","parent_id") VALUES ($1,$2) RETURNING "category_id"`)).
		WithArgs(testCategoryName, nil).
		WillReturnRows(
			sqlmock.NewRows([]string{"category_id"}).
				AddRow(testCategoryId),
		)
	s.mock.ExpectCommit()
	req := model.CreateCategory{
		Title: testCategoryName,
	}
	id, err := s.repository.CreateCategory(s.ctx, req)
	s.NoError(err)
	s.Equal(testCategoryId, id)
}

func (s *CategoryRepositorySuite) Test_Repository_CreateSubcategory() {
	s.mock.ExpectBegin()
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "category" ("title","parent_id") VALUES ($1,$2) RETURNING "category_id"`)).
		WithArgs(testCategoryName, testCategoryId).
		WillReturnRows(
			sqlmock.NewRows([]string{"category_id"}).
				AddRow(testCategoryId),
		)
	s.mock.ExpectCommit()
	req := model.CreateCategory{
		Title:    testCategoryName,
		ParentId: &testCategoryId,
	}

	id, err := s.repository.CreateCategory(s.ctx, req)
	s.NoError(err)
	s.Equal(testCategoryId, id)

}

// func (r CategoryRepository) GetCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
func (s *CategoryRepositorySuite) Test_Repository_GetCategories() {
	pagination := utils.Pagination{
		Limit:      1,
		Page:       1,
		TotalRows:  1,
		TotalPages: 1,
	}
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "category"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"count"}).
				AddRow(1),
		)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "category" LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"category_id", "title", "parent_id"}).
				AddRow(testCategoryId, testCategoryName, nil),
		)
	pag, err := s.repository.GetCategories(s.ctx, pagination)
	pagination.Data = []model.Category{testCategory}
	s.NoError(err)
	s.Equal(pagination, pag)
}

func (s *CategoryRepositorySuite) Test_Repository_GetCategory() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "category" WHERE "category"."category_id" = $1 ORDER BY "category"."category_id" LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"category_id", "title", "parent_id"}).
				AddRow(testCategoryId, testCategoryName, nil),
		)
	cat, err := s.repository.GetCategory(s.ctx, testCategoryId)
	s.NoError(err)
	s.Equal(testCategory, cat)
}

func (s *CategoryRepositorySuite) Test_Repository_DeleteCategory() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "category" WHERE "category"."category_id" = $1`)).
		WithArgs(testCategoryId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteCategory(s.ctx, testCategoryId)
	s.NoError(err)
}

func (s *CategoryRepositorySuite) Test_Repository_CountCategories() {
	testCount := int64(5)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "category" WHERE parent_id = $1`)).
		WithArgs(testCategoryId).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(testCount))

	count, err := s.repository.CountCategories(s.ctx, testCategoryId)
	s.NoError(err)
	s.Equal(testCount, count)
}

func TestCategoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(CategoryRepositorySuite))
}
