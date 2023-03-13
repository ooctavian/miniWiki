package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func uploadProfilePictureHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, header, err := utils.GetImage(r)

		if err != nil {
			logrus.WithContext(r.Context()).Infof("Can't get file: %v", err)
			utils.HandleErrorResponse(w, err)
			return
		}

		req := model.UploadProfilePictureRequest{
			Image:     file,
			ImageName: header.Filename,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.UploadProfilePicture(r.Context(), req)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusCreated, nil)
	}
}
