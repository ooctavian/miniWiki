package controller

import (
	"net/http"
	"os"

	"miniWiki/domain/image/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
)

func uploadResourceImageHandler(service imageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId := chi.URLParam(r, "id")
		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile("File")
		if err != nil {
			utils.Logger.Infof("Error Retrieving the File %v", err)
			utils.Respond(w, http.StatusBadRequest, nil)
			return
		}

		req := model.UploadRequest{
			ImageFolder: os.Getenv("IMAGES_DIR") + "resources",
			ImageName:   resourceId,
			Image:       file,
		}

		err = service.Upload(r.Context(), req)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, nil)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
