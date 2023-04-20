package integrationtests_test

import (
	"net/http"
	"testing"

	"miniWiki/internal/domain/category/model"

	"github.com/stretchr/testify/suite"
)

var (
	testCreateCategory = model.CreateCategory{
		Title: testCategoryTitle,
	}
)

type CategoryCreationSuite struct {
	IntegrationTestSuite
}

func (s *CategoryCreationSuite) TestCategoryCreation() {
	c := s.GetAuthenticatedClient()
	res := c.Post("/categories", testCreateCategory)
	s.Equal(http.StatusCreated, res.StatusCode)
}

func (s *CategoryCreationSuite) TestSubcategoryCreation() {
	c := s.GetAuthenticatedClient()
	res := c.Post("/categories", testCreateCategory)
	s.Equal(http.StatusCreated, res.StatusCode)
	id := s.parseId(res, 2)
	req := model.CreateCategory{
		Title:    testSubcategoryTitle,
		ParentId: &id,
	}

	res = c.Post("/categories", req)
	s.Equal(http.StatusCreated, res.StatusCode)
}

func TestCategoryCreationSuite(t *testing.T) {
	suite.Run(t, new(CategoryCreationSuite))
}
