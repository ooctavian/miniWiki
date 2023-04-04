package service

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"

	"miniWiki/internal/auth/model"
	"miniWiki/internal/auth/query"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
)

func (s *Auth) Login(ctx context.Context, request model.LoginRequest) (*model.SessionResponse, error) {
	acc, err := s.accountQuerier.GetAccount(ctx, request.Account.Email)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, transport.NotFoundError{
			Id:   request.Account.Email,
			Item: "account",
		}
	}
	match, err := s.hash.Equal(request.Account.Password, *acc.Password)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, err
	}

	if !match {
		logrus.WithContext(ctx).Error(err)
		return nil, model.PasswordMismatchError
	}

	sessionID, err := s.generateSessionId(strconv.Itoa(acc.AccountID), request.IpAddress)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, err
	}

	expiresAt := time.Now().Add(30 * time.Minute)
	params := query.CreateSessionParams{
		SessionID: sessionID,
		AccountID: acc.AccountID,
		IpAddress: request.IpAddress,
		UserAgent: request.UserAgent,
		ExpireAt:  expiresAt,
	}

	_, err = s.sessionQuerier.CreateSession(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed creating session: %v", err)
		return nil, err
	}

	if !*acc.Active {
		_, err = s.accountQuerier.UpdateAccountStatus(ctx, true, acc.AccountID)
		if err != nil {
			logrus.WithContext(ctx).Errorf("Failed creating session: %v", err)
			return nil, err
		}
	}

	return &model.SessionResponse{
		SessionId: sessionID,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *Auth) generateSessionId(accId, ipAddress string) (string, error) {
	rawSessionId, _, err := s.hash.GenerateRaw(accId + ipAddress)
	if err != nil {
		return "", err
	}

	sessionID := base64.RawStdEncoding.EncodeToString(rawSessionId)
	return sessionID, nil
}
