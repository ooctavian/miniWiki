package integrationtests_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"

	"github.com/stretchr/testify/suite"
)

type ResourcesRetrievalSuite struct {
	IntegrationTestSuite
}

func (s *ResourcesRetrievalSuite) TestResourcesRetrieval() {
	c := s.GetAuthenticatedClient()
	res := c.Post("/categories", testCreateCategory)
	s.Equal(http.StatusCreated, res.StatusCode)
	res = c.Post("/resources", testCreateResource)
	s.Equal(http.StatusCreated, res.StatusCode)

	res = c.Get("/resources")
	s.Equal(res.StatusCode, 200)
	body, err := json.Marshal(utils.Pagination{
		Limit:      10,
		Page:       1,
		TotalPages: 1,
		TotalRows:  1,
		Data: []model.ResourceResponse{{
			ResourceId:  1,
			Title:       testCreateResource.Title,
			Description: testCreateResource.Description,
			Link:        testCreateResource.Link,
			CategoryId:  &testCreateResource.CategoryId,
			State:       "PUBLIC",
			AuthorId:    1,
		}}})
	s.NoError(err)
	s.JSONEq(c.GetBody(res), string(body))
}

func TestResourcesRetrievalSuite(t *testing.T) {
	suite.Run(t, new(ResourcesRetrievalSuite))
}
