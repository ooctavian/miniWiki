package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// NOTE: should this even be here or should I imitate a cdn ?
func downloadResourceImageHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		req := model.DownloadResourceImageRequest{
			ResourceId: resourceId,
			AccountId:  middleware.GetAccountId(r),
		}

		f, err := service.DownloadResourceImage(r.Context(), req)
		if err != nil {
			logrus.Errorf("Error downloading image %v:", err)
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.RespondWithFile(w, f)
	}
}
