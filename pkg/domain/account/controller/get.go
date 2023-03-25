package controller

import (
	"net/http"

	"miniWiki/pkg/domain/account/model"
	"miniWiki/pkg/middleware"
	"miniWiki/pkg/transport"
)

// swagger:operation GET /account Account getAccount
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

func getAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := model.GetAccountRequest{
			AccountId: middleware.GetAccountId(r),
		}

		res, err := service.GetAccount(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, res)
	}
}
