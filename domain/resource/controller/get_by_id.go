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

func getResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		request := model.GetResourceRequest{
			ResourceId: resourceId,
			AccountId:  middleware.GetAccountId(r),
		}

		resource, err := service.GetResource(r.Context(), request)

		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, resource)
	}
}
