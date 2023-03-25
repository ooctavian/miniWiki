package controller

import (
	"net/http"

	"miniWiki/internal/domain/account/model"
	"miniWiki/internal/middleware"
	"miniWiki/pkg/transport"
)

// swagger:operation DELETE /account Account deactivateAccount
//
// Deactivate account. All posts become private.
//
// ---
// responses:
//   '200':
//     description: 'Account deleted.'
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

func deactivateAccountHandler(service accountService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := model.DeactivateAccountRequest{
			AccountId: middleware.GetAccountId(r),
		}

		err := service.DeactivateAccount(r.Context(), request)
		if err != nil {
			transport.HandleErrorResponse(w, err)
			return
		}

		transport.Respond(w, http.StatusOK, nil)
	}
}
