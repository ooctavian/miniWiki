package model

import "time"

type CreateAccount struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateAccountRequest struct {
	Account CreateAccount
}

type UpdateAccount struct {
	Email    *string `json:"email" validate:"email"`
	Password *string `json:"password"`
}

type UpdateAccountRequest struct {
	Account   UpdateAccount
	AccountId int
}

type GetAccountRequest struct {
	AccountId int
}

type AccountResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}
