package controller

import (
	"net/http"

	"miniWiki/domain/category/model"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func createCategoryHandler(service categoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createCategory := model.CreateCategory{}
		err := utils.Decode(r.Body, &createCategory)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("BadRequest: %v", err)
			return
		}

		request := model.CreateCategoryRequest{
			Category: createCategory,
		}

		err = service.CreateCategory(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusCreated, nil)
	}
}
