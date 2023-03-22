package integrationtests_test

import (
	"net/http"
	"testing"

	"miniWiki/domain/resource/model"

	"github.com/stretchr/testify/suite"
)

type ResourceRetrievalSuite struct {
	IntegrationTestSuite
}

func (s *ResourceRetrievalSuite) TestResourceRetrieval() {
	c := s.GetAuthenticatedClient()
	res := c.Post("/categories", testCreateCategory)
	s.Equal(http.StatusCreated, res.StatusCode)
	id := s.parseId(res, 2)
	testCreateResource.CategoryId = id
	res = c.Post("/resources", testCreateResource)
	s.Equal(http.StatusCreated, res.StatusCode)

	res = c.Get(res.Header.Get("Location"))
	s.Equal(http.StatusOK, res.StatusCode)
	body := s.encode(model.ResourceResponse{
		ResourceId:  1,
		Title:       testCreateResource.Title,
		Description: testCreateResource.Description,
		Link:        testCreateResource.Link,
		CategoryId:  &testCreateResource.CategoryId,
		AuthorId:    1,
	})
	s.JSONEq(c.GetBody(res), body)
}

func TestResourceRetrievalSuite(t *testing.T) {
	suite.Run(t, new(ResourceRetrievalSuite))
}
