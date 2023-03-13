package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func createProfileHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createProfile := model.CreateProfile{}
		err := utils.DecodeJson(r.Body, &createProfile)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.CreateProfileRequest{
			Profile:   createProfile,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.CreateProfile(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusCreated, nil)
	}
}
