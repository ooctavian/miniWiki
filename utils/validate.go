package utils

import (
	"net/url"

	"github.com/go-playground/validator/v10"
)

var (
	validate = NewValidator()
)

func isUrl(fl validator.FieldLevel) bool {
	u, err := url.Parse(fl.Field().String())
	return err == nil && u.Scheme != "" && u.Host != ""
}

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("url", isUrl) //nolint: errcheck // always returns without error
	return v
}
