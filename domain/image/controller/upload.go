package controller

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"miniWiki/domain/image/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func uploadResourceImageHandler(service imageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId := chi.URLParam(r, "id")
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			logrus.WithContext(r.Context()).Error(err)
			utils.HandleErrorResponse(w, err)
			return
		}
		file, _, err := r.FormFile("File")
		if err != nil {
			logrus.Infof("Error Retrieving the File %v", err)
			utils.Respond(w, http.StatusBadRequest, nil)
			return
		}

		contentType := http.DetectContentType(StreamToByte(file))
		if contentType != "image/png" && contentType != "image/jpeg" {
			logrus.Infof("Unsupported image format: %v", contentType)
			utils.ErrorRespond(w, http.StatusBadRequest,
				fmt.Sprintf("Unsupported image format: %v", contentType),
				errors.New("only png and jpeg images are supported"))
			utils.Respond(w, http.StatusBadRequest, nil)
			return
		}

		req := model.UploadRequest{
			ImageFolder: "resources",
			ImageName:   resourceId,
			Image:       file,
		}

		err = service.Upload(r.Context(), req)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		logrus.Error(err)
	}
	return buf.Bytes()
}
