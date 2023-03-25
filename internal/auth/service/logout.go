package service

import (
	"context"

	"miniWiki/internal/auth/model"

	"github.com/sirupsen/logrus"
)

func (s *Auth) Logout(ctx context.Context, request model.LogoutRequest) error {
	_, err := s.sessionQuerier.DeleteSession(ctx, request.SessionId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed deleting session: ")
	}

	return err
}
