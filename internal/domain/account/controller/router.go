package controller

import (
	"context"

	"miniWiki/internal/domain/account/model"

	"github.com/go-chi/chi/v5"
)

type accountService interface {
	CreateAccount(ctx context.Context, request model.CreateAccountRequest) error
	UpdateAccount(ctx context.Context, request model.UpdateAccountRequest) error
	GetAccount(ctx context.Context, request model.GetAccountRequest) (*model.AccountResponse, error)
	GetPublicAccount(ctx context.Context, request model.GetAccountRequest) (*model.PublicAccountResponse, error)
	DeactivateAccount(ctx context.Context, request model.DeactivateAccountRequest) error
	UploadProfilePicture(ctx context.Context, request model.UploadProfilePictureRequest) error
}

func MakeAccountRouter(r chi.Router, service accountService) {
	r.Post("/", createAccountHandler(service))
	r.Get("/{id}", getPublicAccountHandler(service))
}

func MakePrivateAccountRouter(r chi.Router, service accountService) {
	r.Patch("/", updateAccountHandler(service))
	r.Get("/", getAccountHandler(service))
	r.Delete("/", deactivateAccountHandler(service))
	r.Post("/picture", uploadProfilePictureHandler(service))
}
