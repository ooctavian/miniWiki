package controller

import (
	"net/http"

	"miniWiki/domain/auth/model"
	"miniWiki/middleware"
	"miniWiki/utils"

	"github.com/sirupsen/logrus"
)

func logoutHandler(service authService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.Logout(r.Context(), model.LogoutRequest{
			SessionId: middleware.GetSessionId(r),
		})
		if err != nil {
			logrus.WithContext(r.Context()).Errorf("Failed logout %v", err)
			utils.HandleErrorResponse(w, err)
			return
		}

		middleware.LogoutSession(w)
		utils.Respond(w, http.StatusOK, nil)
		return
	}
}
