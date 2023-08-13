package controller

import (
	"net/http"

	"miniWiki/internal/domain/account/model"
	"miniWiki/internal/middleware"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
)

func uploadProfilePictureHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, contentType, header, err := transport.GetImage(r)

		if err != nil {
			logrus.WithContext(r.Context()).Infof("Can't get file: %v", err)
			transport.HandleErrorResponse(w, err)
			return
		}

		req := model.UploadProfilePictureRequest{
			Image:       file,
			ImageName:   header.Filename,
			ContentType: *contentType,
			AccountId:   middleware.GetAccountId(r),
		}

		err = service.UploadProfilePicture(r.Context(), req)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusCreated, nil)
	}
}
