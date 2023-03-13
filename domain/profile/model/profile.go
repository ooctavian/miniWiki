package model

import "io"

type CreateProfile struct {
	Name  string  `json:"name"`
	Alias *string `json:"alias"`
}

type CreateProfileRequest struct {
	Profile   CreateProfile
	AccountId int
}

type UpdateProfile struct {
	Name  *string `json:"name"`
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

type ProfileResponse struct {
	Name       string  `json:"name"`
	Alias      *string `json:"alias,omitempty"`
	ProfileUrl *string `json:"profileUrl,omitempty"`
}
