package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func createResourceHandler(resource resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createResource := model.CreateResource{
			State: "PUBLIC",
		}

		err := utils.DecodeJson(r.Body, &createResource)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.CreateResourceRequest{
			Resource:  createResource,
			AccountId: middleware.GetAccountId(r),
		}

		res, err := resource.CreateResource(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		w.Header().Add("Location", "/resources/"+strconv.Itoa(res.ResourceId))
		utils.Respond(w, http.StatusCreated, nil)
	}
}
