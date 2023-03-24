package controller

import (
	"net/http"

	"miniWiki/domain/auth/model"
	"miniWiki/middleware"
	"miniWiki/transport"

	"github.com/sirupsen/logrus"
)

// swagger:operation POST /refresh Auth refresh
//
// Refersh session token.
//
// ---
// responses:
//   '200':
//     description: 'Token refreshed.'
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

func refreshHandler(service authService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accId := middleware.GetAccountId(r)
		sId := middleware.GetSessionId(r)
		req := model.RefreshRequest{
			AccountId: accId,
			SessionId: sId,
			IpAddress: r.RemoteAddr,
		}

		res, err := service.Refresh(r.Context(), req)
		if err != nil {
			logrus.WithContext(r.Context()).Error(err)
			transport.HandleErrorResponse(w, err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   res.SessionId,
			Expires: res.ExpiresAt,
		})

		transport.Respond(w, http.StatusOK, nil)
	}
}
