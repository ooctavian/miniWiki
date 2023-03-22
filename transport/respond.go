package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

// swagger:model ErrorResponse
type ErrorResponse struct {
	// Status code of response
	Status int `json:"status,omitempty"`
	// Message, usually for the user
	Message string `json:"message,omitempty"`
	// Detail, more detailed message about the problem
	Detail string `json:"detail,omitempty"`
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

	if ok := errors.As(err, &unsupportedContentTypeError{}); ok {
		ErrorRespond(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if ok := errors.As(err, &ForbiddenError{}); ok {
		ErrorRespond(w, http.StatusForbidden, "Not allowed", err)
		return
	}

	if ok := errors.As(err, &schema.MultiError{}); ok {
		ErrorRespond(w, http.StatusBadRequest, "Invalid query parameters", err)
		return
	}

	logrus.Infof("%v", reflect.TypeOf(err).String())
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

func RespondWithFile(w http.ResponseWriter, file io.Reader) {
	buf := streamToByte(file)
	contentType := http.DetectContentType(buf)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", contentType)
	_, err := io.Copy(w, bytes.NewReader(buf))
	if err != nil {
		logrus.Info(err)
		return
	}
}

type NotFoundError struct {
	Item string
	Id   string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Couldn't find %s with id %s", e.Item, e.Id)
}

type ForbiddenError struct{}

func (e ForbiddenError) Error() string {
	return "forbidden"
}

type unsupportedContentTypeError struct {
	contentType string
}

func newUnsupportedContentType(contentType string) unsupportedContentTypeError {
	return unsupportedContentTypeError{
		contentType: contentType,
	}
}

func (e unsupportedContentTypeError) Error() string {
	return fmt.Sprintf("Unsupported content-type : %s", e.contentType)
}
