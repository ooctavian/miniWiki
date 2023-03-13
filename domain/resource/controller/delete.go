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

func deleteResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Infof("Error converting string to int: %v", err)
			return
		}

		req := model.DeleteResourceRequest{
			ResourceId: resourceId,
			AccountId:  middleware.GetAccountId(r),
		}

		err = service.DeleteResource(r.Context(), req)

		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
