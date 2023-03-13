package controller

import (
	"net/http"

	"miniWiki/domain/auth/model"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func loginHandler(service authService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var account model.LoginAccount
		err := utils.DecodeJson(r.Body, &account)
		if err != nil {
			logrus.WithContext(r.Context()).Error(err)
			utils.HandleErrorResponse(w, err)
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
