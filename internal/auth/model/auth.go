package model

import (
	"errors"
	"time"
)

var (
	PasswordMismatchError = errors.New("password mismatch")
)

// swagger:model LoginAccount
type LoginAccount struct {
	// Email of the account
	// example: lorem@example.com
	// required: true
	Email string `json:"email"`
	// Password
	// example: verysecurepassword
	// required: true
	Password string `json:"password"`
}

type LoginRequest struct {
	Account   LoginAccount
	UserAgent string
	IpAddress string
}

type SessionResponse struct {
	SessionId string
}

type RefreshRequest struct {
	AccountId int
	SessionId string
	IpAddress string
}

type LogoutRequest struct {
	SessionId string
}

type Session struct {
	SessionID string
	AccountID int
	IpAddress string
	UserAgent string
	ExpiresAt time.Time
}
