package service

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateCategorySuite struct {
	suite.Suite
}

func (s *CreateCategorySuite) CreateCategorySuccesful() {

}

func TestCreateCategorySuite(t *testing.T) {
	suite.Run(t, new(CreateCategorySuite))
}
