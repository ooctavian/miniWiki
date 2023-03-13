package controller

import (
	"net/http"

	"miniWiki/domain/account/model"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func createAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		account := model.CreateAccount{}

		err := utils.DecodeJson(r.Body, &account)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.CreateAccountRequest{
			Account: account,
		}

		err = service.CreateAccount(r.Context(), request)
		if err != nil {
			utils.HandleErrorResponse(w, err)
			return
		}

		utils.Respond(w, http.StatusCreated, nil)
	}
}
