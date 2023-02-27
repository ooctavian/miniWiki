package integrationtests

import (
	"testing"

	"miniWiki/domain/category/model"

	"github.com/stretchr/testify/suite"
)

type CategoryCreationSuite struct {
	IntegrationTestSuite
}

var (
	testCreateCategory = model.CreateCategory{
		Title:    testCategoryTitle,
		ParentId: nil,
	}

	testCreateSubcategory = model.CreateCategory{
		Title:    testSubcategoryTitle,
		ParentId: &testSubcategoryParentId,
	}
)

func (s *CategoryCreationSuite) TestCategoryCreation() {
	res := s.clt.Post("/categories", testCreateCategory)
	s.Equal(res.StatusCode, 201)
}

func (s *CategoryCreationSuite) TestSubategoryCreation() {
	res := s.clt.Post("/categories", testCreateCategory)
	s.Equal(res.StatusCode, 201)

	res = s.clt.Post("/categories", testCreateSubcategory)
	s.Equal(res.StatusCode, 201)
}

func TestCategoryCreationSuite(t *testing.T) {
	suite.Run(t, new(CategoryCreationSuite))
}
