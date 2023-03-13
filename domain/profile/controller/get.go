package controller

import (
	"net/http"

	"miniWiki/domain/profile/model"
	"miniWiki/middleware"
	"miniWiki/utils"
)

func getProfileHandler(service profileService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := model.GetProfileRequest{
			AccountId: middleware.GetAccountId(r),
		}

		res, err := service.GetProfile(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, res)
	}
}
