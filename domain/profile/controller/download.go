package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/transport"

	"github.com/sirupsen/logrus"
)

// NOTE: should this even be here or should I imitate a cdn ?
func downloadProfilePictureHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.DownloadProfilePictureRequest{
			AccountId: middleware.GetAccountId(r),
		}

		f, err := service.DownloadResourceImage(r.Context(), req)
		if err != nil {
			logrus.Errorf("Error downloading image %v:", err)
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.RespondWithFile(w, f)
	}
}
