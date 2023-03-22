package model

import "io"

// swagger:model CreateProfile
type CreateProfile struct {
	// Name Real name used to show to other people
	// required: true
	Name string `json:"name"`
	// Alias An additional or assumed name
	Alias *string `json:"alias"`
}

type CreateProfileRequest struct {
	Profile   CreateProfile
	AccountId int
}

// swagger:model ProfileResponse
type UpdateProfile struct {
	// Name Real name used to show to other people
	// required: true
	Name *string `json:"name"`
	// Alias An additional or assumed name
	Alias *string `json:"alias"`
}

type UpdateProfileRequest struct {
	Profile   UpdateProfile
	AccountId int
}

type UploadProfilePictureRequest struct {
	AccountId int
	ImageName string
	Image     io.Reader
}

type DownloadProfilePictureRequest struct {
	AccountId int
}

type GetProfileRequest struct {
	AccountId int
}

// swagger:model ProfileResponse
type ProfileResponse struct {
	// Name Real name used to show to other people
	// required: true
	Name string `json:"name"`
	// Alias An additional or assumed name
	Alias *string `json:"alias,omitempty"`
	// PictureUrl path of the profile picture
	PictureUrl *string `json:"pictureUrl,omitempty"`
}
