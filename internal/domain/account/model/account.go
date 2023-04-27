package model

import (
	"io"
	"time"
)

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
	// Real name used to show to other people
	// required: true
	Name string `json:"name" validate:"required"`
	// An additional or assumed name
	Alias string `json:"alias"`
	// path of the profile picture
	PictureUrl string `json:"pictureUrl,omitempty"`
}

func (CreateAccount) TableName() string {
	return "account"
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
	// Real name used to show to other people
	Name *string `json:"name"`
	// Alias An additional or assumed name
	Alias *string `json:"alias"`
	// Status of account
	Active *bool
	// path of the profile picture
	PictureUrl *string `json:"pictureUrl,omitempty"`
}

func (UpdateAccount) TableName() string {
	return "account"
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
	// Status of account
	Active bool `json:"active"`
	// Real name used to show to other people
	Name string `json:"name"`
	// Alias An additional or assumed name
	Alias *string `json:"alias,omitempty"`
	// PictureUrl path of the profile picture
	PictureUrl *string `json:"pictureUrl,omitempty"`
	// CreatedAt the date it was created
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt the last date it was modified
	UpdatedAt time.Time `json:"updatedAt"`
}

// swagger:model PublicAccountResponse
type PublicAccountResponse struct {
	// Real name used to show to other people
	Name string `json:"name"`
	// Alias An additional or assumed name
	Alias *string `json:"alias,omitempty"`
	// path of the profile picture
	PictureUrl *string `json:"pictureUrl,omitempty"`
}

// swagger:model Account
type Account struct {
	ID int `gorm:"column:account_id"`
	// Email of the account
	Email string `json:"email"`
	// Password
	// example: verysecurepassword
	Password string `json:"password"`
	// Status of account
	Active bool `json:"active"`
	// Real name used to show to other people
	Name string `json:"name"`
	// Alias An additional or assumed name
	Alias *string `json:"alias,omitempty"`
	// PictureUrl path of the profile picture
	PictureUrl *string `json:"pictureUrl,omitempty"`
	// CreatedAt the date it was created
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt the last date it was modified
	UpdatedAt time.Time `json:"updatedAt"`
}

type UploadProfilePictureRequest struct {
	AccountId   int
	ImageName   string
	ContentType string
	Image       io.Reader
}

type DownloadProfilePictureRequest struct {
	AccountId int
}
