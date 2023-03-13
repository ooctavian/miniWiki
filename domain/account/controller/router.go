package controller

import (
	"context"

	"miniWiki/domain/account/model"

	"github.com/go-chi/chi/v5"
)

type accountService interface {
	CreateAccount(ctx context.Context, request model.CreateAccountRequest) error
	UpdateAccount(ctx context.Context, request model.UpdateAccountRequest) error
	GetAccount(ctx context.Context, request model.GetAccountRequest) (*model.AccountResponse, error)
}

func MakeAccountRouter(r chi.Router, service accountService) {
	r.Post("/", createAccountHandler(service))
}

func MakePrivateAccountRouter(r chi.Router, service accountService) {
	r.Patch("/", updateAccountHandler(service))
	r.Get("/", getAccountHandler(service))
}
