package service

import (
	"context"

	"miniWiki/internal/auth/model"

	"github.com/sirupsen/logrus"
)

func (s *Auth) Logout(ctx context.Context, request model.LogoutRequest) error {
	err := s.authRepository.DeleteSession(ctx, request.SessionId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed deleting session: ")
	}

	return err
}
