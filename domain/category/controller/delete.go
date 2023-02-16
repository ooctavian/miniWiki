package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/category/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
)

func deleteResourceHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.Logger.WithContext(r.Context()).Infof("Error converting string to int: %v", err)
			return
		}

		req := model.DeleteCategoryRequest{CategoryId: categoryId}

		err = service.DeleteCategory(r.Context(), req)

		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, nil)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
