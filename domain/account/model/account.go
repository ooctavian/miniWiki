package model

import "time"

// swagger:model CreateAccount
type CreateAccount struct {
	// Email of the account
	// example: lorem@example.com
	// required: true
	Email string `json:"email" validate:"required,email"`
	// Password
	// example: verysecurepassword
	// required: true
	Password string `json:"password" validate:"required"`
}

type CreateAccountRequest struct {
	Account CreateAccount
}

// swagger:model UpdateAccount
type UpdateAccount struct {
	// Email of the account
	// example: lorem@example.com
	Email *string `json:"email" validate:"email"`
	// Password
	// example: verysecurepassword
	Password *string `json:"password"`
}

type UpdateAccountRequest struct {
	Account   UpdateAccount
	AccountId int
}

type GetAccountRequest struct {
	AccountId int
}

type DeactivateAccountRequest struct {
	AccountId int
}

// swagger:model AccountResponse
type AccountResponse struct {
	// Email of the account
	Email string `json:"email"`
	// CreatedAt the date it was created
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt the last date it was modified
	UpdateAt time.Time `json:"updateAt"`
}
