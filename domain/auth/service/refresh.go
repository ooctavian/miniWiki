package service

import (
	"context"
	"strconv"
	"time"

	"miniWiki/domain/auth/model"
	"miniWiki/domain/auth/query"

	"github.com/sirupsen/logrus"
)

func (s *Auth) Refresh(ctx context.Context, request model.RefreshRequest) (*model.SessionResponse, error) {
	sId, err := s.generateSessionId(strconv.Itoa(request.AccountId), request.IpAddress)
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed updating session: %v", err)
		return nil, err
	}

	expiresAt := time.Now().Add(30 * time.Minute)

	_, err = s.sessionQuerier.UpdateSessionID(ctx,
		query.UpdateSessionIDParams{
			NewSessionID: sId,
			ExpireAt:     expiresAt,
			OldSessionID: request.SessionId,
		})

	if err != nil {
		logrus.WithContext(ctx).Infof("Failed updating session: %v", err)
		return nil, err
	}

	return &model.SessionResponse{
		SessionId: sId,
		ExpiresAt: expiresAt,
	}, nil
}
