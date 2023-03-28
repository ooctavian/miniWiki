package security

import (
	"errors"
	"regexp"
)

var (
	InvalidPasswordUpper   = errors.New("password should have at least one upper character")
	InvalidPasswordLower   = errors.New("password should have at least one lower character")
	InvalidPasswordDigit   = errors.New("password should have at least one digit")
	InvalidPasswordSpecial = errors.New("password should have at least one special character")
	InvalidPasswordLength  = errors.New("password should be at least 8 characters long")
)

func ValidatePassword(password []byte) error {
	if len(password) < 8 {
		return InvalidPasswordLength
	}
	re := regexp.MustCompile("[A-Z]")
	if !re.Match(password) {
		return InvalidPasswordUpper
	}
	re = regexp.MustCompile("[a-z]")
	if !re.Match(password) {
		return InvalidPasswordLower
	}
	re = regexp.MustCompile("[0-9]")
	if !re.Match(password) {
		return InvalidPasswordDigit
	}
	re = regexp.MustCompile(`\W`)
	if !re.Match(password) {
		return InvalidPasswordSpecial
	}
	return nil
}
