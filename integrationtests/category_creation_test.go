package integrationtests

import (
	"net/http"
	"testing"

	"miniWiki/domain/category/model"

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

	req := model.CreateCategory{
		Title:    testSubcategoryTitle,
		ParentId: testSubcategoryParentId,
	}

	res = c.Post("/categories", req)
	s.Equal(http.StatusCreated, res.StatusCode)
}

func TestCategoryCreationSuite(t *testing.T) {
	suite.Run(t, new(CategoryCreationSuite))
}
