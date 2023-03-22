package model

import (
	"io"
	"time"

	"miniWiki/domain/category/model"
)

// ResourceResponse Resource information
// swagger:model ResourceResponse
type ResourceResponse struct {
	// Id of resource
	// example: 1
	ResourceId int `json:"resourceId"`
	// Title of resource
	// example: Lorem ipsum
	Title string `json:"title"`
	// Description of resource
	// example: Lorem Ipsum is simply dummy text of the printing and typesetting industry.
	Description string `json:"description,omitempty"`
	// Link of resource
	// example: Lorem Ipsum is simply dummy text of the printing and typesetting industry.
	Link string `json:"link"`
	// State of resource, can be either PUBLIC or PRIVATE
	// example: PUBLIC
	State string `json:"state"`
	// CategoryId ID of the category that the resource is a part of
	// example: 1
	CategoryId *int `json:"categoryId,omitempty"`
	// AuthorId ID of resource's author
	// example: 1
	AuthorId int `json:"authorId,omitempty"`
}

type DeleteResourceRequest struct {
	ResourceId int
	AccountId  int
}

type GetResourceRequest struct {
	ResourceId int
	AccountId  int
}

type GetResourcesRequest struct {
	Filters   GetResourcesFilters
	AccountId int
}

// swagger:model GetResourcesFilters
type GetResourcesFilters struct {
	// Title
	Title string `schema:"title"`
	// Link
	Link string `schema:"link"`
	// Categories
	Categories []int `schema:"categories"`
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

// swagger:model UpdateResource
type UpdateResource struct {
	// Title of resource
	// example: Lorem ipsum
	Title *string `json:"title"`
	// Description of resource
	// example: Lorem Ipsum is simply dummy text of the printing and typesetting industry.
	Description *string `json:"description"`
	// Link of resource
	// example: Lorem Ipsum is simply dummy text of the printing and typesetting industry.
	Link *string `json:"link" validate:"url"`
	// Id of the category that the resource is a part of
	// example: 1
	CategoryId *int `json:"categoryId"`
	// State of resource, can be either PUBLIC or PRIVATE
	// example: PUBLIC
	State *string `json:"state"`
}

type CreateResourceRequest struct {
	Resource  CreateResource
	AccountId int
}

// swagger:model CreateResource
type CreateResource struct {
	// Title of resource
	// example: Lorem ipsum
	// required: true
	Title string `json:"title" validate:"required"`
	// Description of resource
	// example: Lorem Ipsum is simply dummy text of the printing and typesetting industry.
	Description string `json:"description"`
	// Link of resource
	// example: Lorem Ipsum is simply dummy text of the printing and typesetting industry.
	// required: true
	Link string `json:"link" validate:"required,url"`
	// Id of the category that the resource is a part of
	// example: 1
	CategoryId int `json:"categoryId"`
	// State of resource, can be either PUBLIC or PRIVATE
	// example: PUBLIC
	State string `json:"state"`
}

type Resource struct {
	ID          uint `gorm:"column:category_id"`
	Title       string
	Description string
	Link        string
	State       string
	Image       string
	AuthorId    uint
	CategoryId  *uint
	Category    model.Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
