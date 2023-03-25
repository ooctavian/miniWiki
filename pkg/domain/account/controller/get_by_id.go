package controller

import (
	"net/http"
	"strconv"

	"miniWiki/pkg/domain/account/model"
	"miniWiki/pkg/transport"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// swagger:operation GET /user/{id} Account getAccount
//
// Get details of current logged account.
//
// ---
// responses:
//   '200':
//     description: 'Account info'
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/AccountResponse"
//   '400':
//     description: Invalid body request.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '401':
//     description: Unauthorized.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func getPublicAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accountId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Error converting string to int: %v", err)
			return
		}
		request := model.GetAccountRequest{
			AccountId: accountId,
		}

		res, err := service.GetPublicAccount(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, res)
	}
}
