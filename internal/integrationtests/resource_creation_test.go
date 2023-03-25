package integrationtests_test

import (
	"net/http"
	"testing"

	"miniWiki/internal/domain/resource/model"

	"github.com/stretchr/testify/suite"
)

var (
	testCreateResource = model.CreateResource{
		Title:       testResourceTitle,
		Description: testResourceDescription,
		Link:        testResourceLink,
		CategoryId:  testResourceCategoryId,
		State:       testResourceState,
	}
)

type ResourceCreationSuite struct {
	IntegrationTestSuite
}

func (s *ResourceCreationSuite) TestResourceCreation() {
	c := s.GetAuthenticatedClient()
	res := c.Post("/categories", testCreateCategory)
	s.Equal(http.StatusCreated, res.StatusCode)
	res = c.Post("/resources", testCreateResource)
	s.Equal(http.StatusCreated, res.StatusCode)
}

func TestResourceCreationSuite(t *testing.T) {
	suite.Run(t, new(ResourceCreationSuite))
}
