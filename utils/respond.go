package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

func Respond(w http.ResponseWriter, code int, v interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if v == nil {
		return
	}
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		logrus.Info("Encoding error: %v", err)
		return
	}
}

func HandleErrorResponse(w http.ResponseWriter, err error) {
	if ok := errors.As(err, &NotFoundError{}); ok {
		ErrorRespond(w, http.StatusNotFound, "Not Found", err)
		return
	}

	var invalidUnmarshalErr *json.InvalidUnmarshalError
	if ok := errors.As(err, &invalidUnmarshalErr); ok {
		ErrorRespond(w, http.StatusBadRequest, "Invalid body request", err)
		return
	}

	var unmarshalTypeErr *json.UnmarshalTypeError
	if ok := errors.As(err, &unmarshalTypeErr); ok {
		ErrorRespond(w, http.StatusBadRequest, "Invalid body request", err)
		return
	}

	var syntaxErr *json.SyntaxError
	if ok := errors.As(err, &syntaxErr); ok {
		ErrorRespond(w, http.StatusBadRequest, "Invalid body request", err)
		return
	}

	if ok := errors.As(err, &validator.ValidationErrors{}); ok {
		ErrorRespond(w, http.StatusBadRequest, "Invalid body request", err)
		return
	}

	if ok := errors.As(err, &schema.MultiError{}); ok {
		ErrorRespond(w, http.StatusBadRequest, "Invalid query parameters", err)
		return
	}

	logrus.Info("%v", reflect.TypeOf(err).String())
	ErrorRespond(w, http.StatusInternalServerError, "Internal Server Error", err)
}

func ErrorRespond(w http.ResponseWriter, code int, message string, err error) {
	response := ErrorResponse{
		Status:  code,
		Message: message,
		Detail:  err.Error(),
	}
	Respond(w, code, response)
}

type NotFoundError struct {
	Item string
	Id   string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Couldn't find %s with id %s", e.Item, e.Id)
}
