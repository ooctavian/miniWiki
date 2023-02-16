package controller

import (
	"net/http"

	"miniWiki/utils"
)

func getResourcesHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resources, err := service.GetResources(r.Context())
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, nil)
			return
		}

		utils.Respond(w, http.StatusOK, resources)
	}
}
