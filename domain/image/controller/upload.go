package controller

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"miniWiki/domain/image/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func uploadResourceImageHandler(service imageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId := chi.URLParam(r, "id")
		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile("File")
		if err != nil {
			logrus.Infof("Error Retrieving the File %v", err)
			utils.Respond(w, http.StatusBadRequest, nil)
			return
		}

		contentType := http.DetectContentType(StreamToByte(file))
		if contentType != "image/png" || contentType != "image/jpeg" {
			logrus.Infof("Unsupported image format %v", contentType)
			utils.Respond(w, http.StatusBadRequest, nil)
		}

		req := model.UploadRequest{
			ImageFolder: os.Getenv("IMAGES_DIR") + "resources",
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
	buf.ReadFrom(stream)
	return buf.Bytes()
}
