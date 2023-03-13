package model

import "time"

type LoginAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Account   LoginAccount
	UserAgent string
	IpAddress string
}

type SessionResponse struct {
	SessionId string    `json:"sessionId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type RefreshRequest struct {
	AccountId int
	SessionId string
	IpAddress string
}

type LogoutRequest struct {
	SessionId string
}
