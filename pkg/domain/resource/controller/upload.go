package controller

import (
	"net/http"
	"strconv"

	"miniWiki/pkg/domain/resource/model"
	"miniWiki/pkg/middleware"
	transport2 "miniWiki/pkg/transport"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func uploadResourceImageHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		file, header, err := transport2.GetImage(r)

		if err != nil {
			logrus.WithContext(r.Context()).Infof("Can't get file: %v", err)
			transport2.HandleErrorResponse(w, err)
			return
		}

		req := model.UploadResourceImageRequest{
			ResourceId: resourceId,
			Image:      file,
			ImageName:  header.Filename,
			AccountId:  middleware.GetAccountId(r),
		}

		err = service.UploadResourceImage(r.Context(), req)
		if err != nil {
			transport2.HandleErrorResponse(w, err)
			return
		}

		transport2.Respond(w, http.StatusCreated, nil)
	}
}
