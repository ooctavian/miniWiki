package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
)

func getResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.Logger.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}
		request := model.GetResourceRequest{
			ResourceId: resourceId,
		}

		resource, err := service.GetResource(r.Context(), request)

		if err != nil {
			utils.Respond(w, http.StatusNotFound, nil)
			return
		}

		utils.Respond(w, http.StatusOK, resource)
	}
}
