package controller

import (
	"net/http"

	"miniWiki/domain/account/model"
	"miniWiki/transport"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

// swagger:operation POST /account Account createAccount
//
// Create a Account.
//
// ---
// parameters:
// - name: 'Account'
//   in: body
//   schema:
//     "$ref": '#/definitions/CreateAccount'
// responses:
//   '201':
//     description: 'Account created.'
//   '400':
//     description: Invalid body request.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func createAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		account := model.CreateAccount{}

		err := utils.DecodeJson(r.Body, &account)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.CreateAccountRequest{
			Account: account,
		}

		err = service.CreateAccount(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusCreated, nil)
	}
}
