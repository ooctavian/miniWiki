package controller

import (
	"net/http"

	"miniWiki/domain/resource/model"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func createResourceHandler(resource resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createResource := model.CreateResource{}
		err := utils.Decode(r.Body, &createResource)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.CreateResourceRequest{
			Resource: createResource,
		}

		err = resource.CreateResource(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusCreated, nil)
	}
}
