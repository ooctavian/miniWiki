package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
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
	var notFoundErr *NotFoundError
	if ok := errors.As(err, &notFoundErr); ok {
		ErrorRespond(w, http.StatusNotFound, "Not Found", err)
		return
	}

	ErrorRespond(w, http.StatusInternalServerError, "Internal Server Error", err)
}

func ErrorRespond(w http.ResponseWriter, code int, message string, err error) {
	response := ErrorResponse{
		Code:    code,
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
