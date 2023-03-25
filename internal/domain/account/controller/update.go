package controller

import (
	"net/http"

	"miniWiki/internal/domain/account/model"
	"miniWiki/internal/middleware"
	"miniWiki/pkg/transport"
	"miniWiki/pkg/utils"

	"github.com/sirupsen/logrus"
)

// swagger:operation PATCH /account Account updateAccount
//
// Update the current logged in account.
//
// ---
// parameters:
// - name: 'Account'
//   in: body
//   schema:
//     "$ref": '#/definitions/UpdateAccount'
// responses:
//   '204':
//     description: 'Account updated.'
//   '400':
//     description: Invalid body request.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '401':
//     description: Unauthorized.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '403':
//     description: Forbidden.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func updateAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		account := model.UpdateAccount{}

		err := utils.DecodeJson(r.Body, &account)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			logrus.WithContext(r.Context()).Infof("Invalid body request: %v", err)
			return
		}

		request := model.UpdateAccountRequest{
			Account:   account,
			AccountId: middleware.GetAccountId(r),
		}

		err = service.UpdateAccount(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, nil)
	}
}
