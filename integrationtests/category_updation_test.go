package integrationtests

import (
	"net/http"
	"testing"

	"miniWiki/domain/category/model"

	"github.com/stretchr/testify/suite"
)

type CategoryUpdationSuite struct {
	IntegrationTestSuite
}

func (s *CategoryUpdationSuite) TestCategoryUpdation() {
	c := s.GetAuthenticatedClient()
	res := c.Post("/categories", testCreateCategory)
	s.Equal(res.StatusCode, http.StatusCreated)
	req := model.UpdateCategory{
		Title: "Updated",
	}
	res = c.Patch("/categories/1", req)
	s.Equal(res.StatusCode, http.StatusNoContent)
}

func TestCategoryUpdationSuite(t *testing.T) {
	suite.Run(t, new(CategoryUpdationSuite))
}
