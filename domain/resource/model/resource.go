package model

import "io"

type ResourceResponse struct {
	ResourceId  int    `json:"resourceId"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link"`
	CategoryId  *int   `json:"categoryId,omitempty"`
}

type DeleteResourceRequest struct {
	ResourceId int `json:"resourceId"`
}

type GetResourceRequest struct {
	ResourceId int `json:"resourceId"`
}

type UpdateResourceRequest struct {
	Resource   UpdateResource
	ResourceId int
}

type UploadResourceImageRequest struct {
	ResourceId string
	Image      io.Reader
}

type UpdateResource struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link" validate:"url"`
	CategoryId  int    `json:"categoryId"`
}

type CreateResourceRequest struct {
	Resource CreateCategory
}

type CreateCategory struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Link        string `json:"link" validate:"required,url"`
	CategoryId  int    `json:"categoryId"`
}
