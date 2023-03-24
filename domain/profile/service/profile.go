package service

import (
	iService "miniWiki/domain/image/service"
	"miniWiki/domain/profile/query"
)

type Profile struct {
	profileQuerier query.Querier
	imageService   iService.ImageService
}

func NewProfile(profileQuerier query.Querier, imageService iService.ImageService) *Profile {
	return &Profile{
		profileQuerier: profileQuerier,
		imageService:   imageService,
	}
}
