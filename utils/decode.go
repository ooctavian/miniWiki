package utils

import (
	"encoding/json"
	"io"
)

func Decode(body io.ReadCloser, v interface{}) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}

	return validate.Struct(v)
}
