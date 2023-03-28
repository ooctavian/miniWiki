package security

import (
	"errors"
	"regexp"
)

var (
	InvalidPasswordUpperErr   = errors.New("password should have at least one upper character")
	InvalidPasswordLowerErr   = errors.New("password should have at least one lower character")
	InvalidPasswordDigitErr   = errors.New("password should have at least one digit")
	InvalidPasswordSpecialErr = errors.New("password should have at least one special character")
	InvalidPasswordLengthErr  = errors.New("password should be at least 8 characters long")
)

func ValidatePassword(password []byte) error {
	if len(password) < 8 {
		return InvalidPasswordLengthErr
	}
	re := regexp.MustCompile("[A-Z]")
	if !re.Match(password) {
		return InvalidPasswordUpperErr
	}
	re = regexp.MustCompile("[a-z]")
	if !re.Match(password) {
		return InvalidPasswordLowerErr
	}
	re = regexp.MustCompile("[0-9]")
	if !re.Match(password) {
		return InvalidPasswordDigitErr
	}
	re = regexp.MustCompile(`\W`)
	if !re.Match(password) {
		return InvalidPasswordSpecialErr
	}
	return nil
}
