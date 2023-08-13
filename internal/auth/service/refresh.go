package service

import (
	"context"
	"strconv"
	"time"

	"miniWiki/internal/auth/model"

	"github.com/sirupsen/logrus"
)

func (s *Auth) Refresh(ctx context.Context, request model.RefreshRequest) (*model.SessionResponse, error) {
	sId, err := s.generateSessionId(strconv.Itoa(request.AccountId), request.IpAddress)
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed updating session: %v", err)
		return nil, err
	}

	expiresAt := time.Now().Add(s.sessionDuration)
	err = s.authRepository.UpdateSession(ctx, request.SessionId,
		model.Session{
			ExpireAt:  expiresAt,
			SessionID: sId,
		})

	if err != nil {
		logrus.WithContext(ctx).Infof("Failed updating session: %v", err)
		return nil, err
	}

	return &model.SessionResponse{
		SessionId: sId,
	}, nil
}
