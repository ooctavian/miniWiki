package controller

import (
	"net/http"

	"miniWiki/domain/account/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func updateAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		account := model.UpdateAccount{}

		err := utils.DecodeJson(r.Body, &account)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.UpdateAccountRequest{
			Account:   account,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.UpdateAccount(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusOK, nil)
	}
}
