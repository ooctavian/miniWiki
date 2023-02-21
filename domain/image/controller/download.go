package controller

import (
	"io"
	"net/http"
	"os"

	"miniWiki/domain/image/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func downloadResourceImageHandler(service imageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId := chi.URLParam(r, "id")

		request := model.DownloadRequest{
			ImageFolder: os.Getenv("IMAGES_DIR") + "resources",
			ImageName:   resourceId,
		}

		f, err := service.Download(r.Context(), request)
		if err != nil {
			logrus.Errorf("Error downloading image %v:", err)
			utils.HandleErrorResponse(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "image/png")
		io.Copy(w, f)
	}
}
