package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func updateProfileHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateProfile := model.UpdateProfile{}
		err := utils.DecodeJson(r.Body, &updateProfile)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.UpdateProfileRequest{
			Profile:   updateProfile,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.UpdateProfile(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
