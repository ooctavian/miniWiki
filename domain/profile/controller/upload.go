package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/transport"

	"github.com/sirupsen/logrus"
)

func uploadProfilePictureHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, header, err := transport.GetImage(r)

		if err != nil {
			logrus.WithContext(r.Context()).Infof("Can't get file: %v", err)
			transport.HandleErrorResponse(w, err)
			return
		}

		req := model.UploadProfilePictureRequest{
			Image:     file,
			ImageName: header.Filename,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.UploadProfilePicture(r.Context(), req)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusCreated, nil)
	}
}
