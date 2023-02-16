package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
)

func deleteResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.Logger.WithContext(r.Context()).Infof("Error converting string to int: %v", err)
			return
		}

		req := model.DeleteResourceRequest{ResourceId: resourceId}

		err = service.DeleteResource(r.Context(), req)

		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, nil)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
