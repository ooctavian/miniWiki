package controller

import (
	"net/http"

	"miniWiki/domain/account/model"
	"miniWiki/middleware"
	"miniWiki/utils"
)

func getAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := model.GetAccountRequest{
			AccountId: middleware.GetAccountId(r),
		}

		res, err := service.GetAccount(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, res)
	}
}
