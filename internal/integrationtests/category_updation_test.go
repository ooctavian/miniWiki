package integrationtests_test

import (
	"net/http"
	"testing"

	"miniWiki/internal/domain/category/model"

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
	res = c.Patch(res.Header.Get("Location"), req)
	s.Equal(res.StatusCode, http.StatusNoContent)
}

func TestCategoryUpdationSuite(t *testing.T) {
	suite.Run(t, new(CategoryUpdationSuite))
}
