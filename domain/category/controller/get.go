package controller

import (
	"net/http"

	"miniWiki/domain/category/model"
	"miniWiki/middleware"
	"miniWiki/utils"
)

func getResourcesHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.GetCategoriesRequest{
			AccountId: middleware.GetAccountId(r),
		}

		categories, err := service.GetCategories(r.Context(), req)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, categories)
	}
}
