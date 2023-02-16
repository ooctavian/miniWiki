package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/category/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
)

func updateResourceHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateCategory := model.UpdateCategory{}
		err := utils.Decode(r.Body, &updateCategory)
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

		request := model.UpdateCategoryRequest{
			Category:   updateCategory,
			CategoryId: resourceId,
		}

		err = service.UpdateCategory(r.Context(), request)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, nil)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
