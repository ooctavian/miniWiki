package security

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PasswordValidationTestSuite struct {
	suite.Suite
}

func (s *PasswordValidationTestSuite) TestPasswordLength() {
	err := ValidatePassword([]byte("pass"))
	s.Error(err)
	s.Equal(InvalidPasswordLengthErr, err)
}

func (s *PasswordValidationTestSuite) TestPasswordUpperChar() {
	err := ValidatePassword([]byte("pas$w0rd"))
	s.Error(err)
	s.Equal(InvalidPasswordUpperErr, err)
}

func (s *PasswordValidationTestSuite) TestPasswordLowerChar() {
	err := ValidatePassword([]byte("PAS$W0RD"))
	s.Error(err)
	s.Equal(InvalidPasswordLowerErr, err)
}

func (s *PasswordValidationTestSuite) TestPasswordDigit() {
	err := ValidatePassword([]byte("PAS$WOrd"))
	s.Error(err)
	s.Equal(InvalidPasswordDigitErr, err)
}

func (s *PasswordValidationTestSuite) TestPasswordSpecial() {
	err := ValidatePassword([]byte("PASSW0rd"))
	s.Error(err)
	s.Equal(InvalidPasswordSpecialErr, err)
}

func TestPasswordValidationTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordValidationTestSuite))
}
