package model

import "time"

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