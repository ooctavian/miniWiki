package controller

import (
	"context"

	"miniWiki/pkg/domain/auth/model"

	"github.com/go-chi/chi/v5"
)

type authService interface {
	Login(ctx context.Context, request model.LoginRequest) (*model.SessionResponse, error)
	Refresh(ctx context.Context, request model.RefreshRequest) (*model.SessionResponse, error)
	Logout(ctx context.Context, request model.LogoutRequest) error
}

func MakeAuthRouter(r chi.Router, service authService) {
	r.Post("/login", loginHandler(service))
}

func MakeProtectedAuthRouter(r chi.Router, service authService) {
	r.Post("/logout", logoutHandler(service))
}
