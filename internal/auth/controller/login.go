package controller

import (
	"net/http"

	"miniWiki/internal/auth/model"
	"miniWiki/internal/middleware"
	"miniWiki/pkg/transport"
	"miniWiki/pkg/utils"

	"github.com/sirupsen/logrus"
)

// swagger:operation POST /login Auth login
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
//   '201':
//     description: 'Authenticated.'
//     headers:
//       Set-Cookie:
//         type: string
//         description: A cookie with session id.
//         example: session_id=abcde12345; Path=/; HttpOnly
//   '400':
//     description: Invalid body request.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/definitions/ErrorResponse"

func loginHandler(service authService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var account model.LoginAccount
		err := utils.DecodeJson(r.Body, &account)
		if err != nil {
			logrus.WithContext(r.Context()).Error(err)
			transport.HandleErrorResponse(w, err)
			return
		}

		req := model.LoginRequest{
			Account:   account,
			UserAgent: r.UserAgent(),
			IpAddress: r.RemoteAddr,
		}

		res, err := service.Login(r.Context(), req)
		if err != nil {
			logrus.WithContext(r.Context()).Error(err)
			transport.HandleErrorResponse(w, err)
			return
		}

		middleware.SetSessionCookie(w, res.SessionId)

		transport.Respond(w, http.StatusOK, nil)
	}
}
