package utils

import (
	"encoding/json"
	"io"

	"github.com/gorilla/schema"
)

var (
	schemaDecoder = schema.NewDecoder()
)

func DecodeJson(body io.ReadCloser, v interface{}) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}

	return validate.Struct(v)
}

func DecodeQuery(v interface{}, query map[string][]string) error {
	err := schemaDecoder.Decode(v, query)
	if err != nil {
		return err
	}

	return nil
}
