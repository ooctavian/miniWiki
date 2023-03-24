package controller

import (
	"context"
	"io"

	"miniWiki/domain/profile/model"

	"github.com/go-chi/chi/v5"
)

type profileService interface {
	CreateProfile(ctx context.Context, request model.CreateProfileRequest) error
	DownloadResourceImage(ctx context.Context, request model.DownloadProfilePictureRequest) (io.Reader, error)
	UpdateProfile(ctx context.Context, request model.UpdateProfileRequest) error
	UploadProfilePicture(ctx context.Context, request model.UploadProfilePictureRequest) error
	GetProfile(ctx context.Context, request model.GetProfileRequest) (*model.ProfileResponse, error)
}

func MakeProfileRouter(r chi.Router, service profileService) {
	r.Post("/", createProfileHandler(service))
	r.Get("/", getProfileHandler(service))
	r.Patch("/", updateProfileHandler(service))
	r.Post("/picture", uploadProfilePictureHandler(service))
	r.Get("/picture", downloadProfilePictureHandler(service))
}

func MakeProfilesRouter(r chi.Router, service profileService) {
	r.Get("/{id}", getProfileByIdHandler(service))
}
