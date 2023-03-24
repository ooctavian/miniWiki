package controller

import (
	"net/http"

	"miniWiki/domain/auth/model"
	"miniWiki/middleware"
	"miniWiki/transport"

	"github.com/sirupsen/logrus"
)

// swagger:operation POST /logout Auth logout
//
// Login into an existing account.
//
// ---
// parameters:
// - name: 'Login'
//   in: body
//   schema:
//     "$ref": '#/definitions/LoginAccount'
// responses:
//   '200':
//     description: 'Succesfully logged out.'
//     headers:
//       Set-Cookie:
//         type: string
//         description: A cookie with session id.
//         example: session_id=""; Path=/; HttpOnly
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func logoutHandler(service authService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.Logout(r.Context(), model.LogoutRequest{
			SessionId: middleware.GetSessionId(r),
		})
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Failed logout %v", err)
			transport.HandleErrorResponse(w, err)
			return
		}

		middleware.LogoutSession(w)
		transport.Respond(w, http.StatusOK, nil)
	}
}
