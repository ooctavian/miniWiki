package service

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"

	"miniWiki/internal/auth/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
)

func (s *Auth) Login(ctx context.Context, request model.LoginRequest) (*model.SessionResponse, error) {
	acc, err := s.accountRepository.GetAccountByEmail(ctx, request.Account.Email)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, transport.NotFoundError{
			Id:   request.Account.Email,
			Item: "account",
		}
	}
	match, err := s.hash.Equal(request.Account.Password, acc.Password)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, err
	}

	if !match {
		logrus.WithContext(ctx).Error(err)
		return nil, model.PasswordMismatchError
	}

	sessionID, err := s.generateSessionId(strconv.Itoa(acc.ID), request.IpAddress)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nil, err
	}

	expiresAt := time.Now().Add(s.sessionDuration)
	session := model.Session{
		SessionID: sessionID,
		AccountID: acc.ID,
		IpAddress: request.IpAddress,
		UserAgent: request.UserAgent,
		ExpiresAt: expiresAt,
	}

	err = s.authRepository.CreateSession(ctx, session)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed creating session: %v", err)
		return nil, err
	}

	if !acc.Active {
		err = s.accountRepository.UpdateAccountStatus(ctx, acc.ID, true)
		if err != nil {
			logrus.WithContext(ctx).Errorf("Failed creating session: %v", err)
			return nil, err
		}
	}

	return &model.SessionResponse{
		SessionId: sessionID,
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
