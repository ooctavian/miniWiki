package model

import "io"

type ResourceResponse struct {
	ResourceId  int    `json:"resourceId"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link"`
	CategoryId  *int   `json:"categoryId,omitempty"`
	AuthorId    int    `json:"authorId"`
}

type DeleteResourceRequest struct {
	ResourceId int `json:"resourceId"`
	AccountId  int
}

type GetResourceRequest struct {
	ResourceId int `json:"resourceId"`
	AccountId  int
}

type GetResourcesRequest struct {
	Filters   GetResourcesFilters
	AccountId int
}

type GetResourcesFilters struct {
	Title      string `schema:"title"`
	Link       string `schema:"link"`
	Categories []int  `schema:"categories"`
}

type UpdateResourceRequest struct {
	Resource   UpdateResource
	ResourceId int
	AccountId  int
}

type DownloadResourceImageRequest struct {
	ResourceId int
	AccountId  int
}

type UploadResourceImageRequest struct {
	ResourceId int
	AccountId  int
	ImageName  string
	Image      io.Reader
}

type UpdateResource struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Link        *string `json:"link" validate:"url"`
	CategoryId  *int    `json:"categoryId"`
	State       *string `json:"state"`
}

type CreateResourceRequest struct {
	Resource  CreateResource
	AccountId int
}

type CreateResource struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Link        string `json:"link" validate:"required,url"`
	CategoryId  int    `json:"categoryId"`
	State       string `json:"state"`
}
