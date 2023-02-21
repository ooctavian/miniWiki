package controller

import (
	"net/http"

	"miniWiki/utils"
)

func getResourcesHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := service.GetCategories(r.Context())
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, categories)
	}
}
