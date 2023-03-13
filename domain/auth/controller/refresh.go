package controller

import (
	"net/http"

	"miniWiki/domain/auth/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

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
			utils.HandleErrorResponse(w, err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   res.SessionId,
			Expires: res.ExpiresAt,
		})

		utils.Respond(w, http.StatusOK, nil)
		return
	}
}
