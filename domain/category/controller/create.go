package controller

import (
	"net/http"

	"miniWiki/domain/category/model"
	"miniWiki/utils"
)

func createCategoryHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createCategory := model.CreateCategory{}
		err := utils.Decode(r.Body, &createCategory)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, nil)
			utils.Logger.WithContext(r.Context()).Infof("BadRequest: %v", err)
			return
		}

		request := model.CreateCategoryRequest{
			Category: createCategory,
		}

		err = service.CreateCategory(r.Context(), request)
		if err != nil {
			utils.Respond(w, http.StatusInternalServerError, nil)
			return
		}

		utils.Respond(w, http.StatusCreated, nil)
	}
}
