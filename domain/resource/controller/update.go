package controller

import (
	"net/http"
	"strconv"

	"miniWiki/domain/resource/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func updateResourceHandler(service resourceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateResource := model.UpdateResource{}
		err := utils.DecodeJson(r.Body, &updateResource)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, nil)
			logrus.WithContext(r.Context()).Infof("BadRequest: %v", err)
			return
		}

		resourceId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}

		request := model.UpdateResourceRequest{
			Resource:   updateResource,
			ResourceId: resourceId,
			AccountId:  middleware.GetAccountId(r),
		}

		res, err := service.UpdateResource(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		w.Header().Add("Location", "/resources/"+strconv.Itoa(res.ResourceId))
		utils.Respond(w, http.StatusOK, nil)
	}
}
