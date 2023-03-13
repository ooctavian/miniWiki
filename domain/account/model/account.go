package model

type CreateAccount struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateAccountRequest struct {
	Account CreateAccount
}
