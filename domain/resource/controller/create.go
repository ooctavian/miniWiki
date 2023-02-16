package controller

import (
	"net/http"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"
)

func createResourceHandler(resource resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createResource := model.CreateCategory{}
		err := utils.Decode(r.Body, &createResource)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, nil)
			utils.Logger.WithContext(r.Context()).Infof("BadRequest: %v", err)
			return
		}

		request := model.CreateResourceRequest{
			Resource: createResource,
		}

		err = resource.CreateResource(r.Context(), request)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, nil)
			return
		}

		utils.Respond(w, http.StatusCreated, nil)
	}
}
