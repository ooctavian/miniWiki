package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/category/model"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func getResourceHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		request := model.GetCategoryRequest{
			CategoryId: categoryId,
		}
		category, err := service.GetCategory(r.Context(), request)

		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, category)
	}
}
