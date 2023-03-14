package controller

import (
	"net/http"

	"miniWiki/domain/account/model"
	"miniWiki/middleware"
	"miniWiki/utils"
)

func deactivateAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := model.DeactivateAccountRequest{
			AccountId: middleware.GetAccountId(r),
		}

		err := service.DeactivateAccount(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
