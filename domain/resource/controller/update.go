package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
)

func updateResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateResource := model.UpdateResource{}
		err := utils.Decode(r.Body, &updateResource)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, nil)
			utils.Logger.WithContext(r.Context()).Infof("BadRequest: %v", err)
			return
		}

		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.Logger.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		request := model.UpdateResourceRequest{
			Resource:   updateResource,
			ResourceId: resourceId,
		}

		err = service.UpdateResource(r.Context(), request)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, nil)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
