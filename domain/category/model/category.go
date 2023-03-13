package model

type CreateCategory struct {
	Title    string `json:"title" validate:"required"`
	ParentId int    `json:"parentId"`
}

type CreateCategoryRequest struct {
	Category  CreateCategory
	AccountId int
}

type UpdateCategory struct {
	Title    string `json:"title"`
	ParentId int    `json:"parentId"`
}

type UpdateCategoryRequest struct {
	CategoryId int
	AccountId  int
	Category   UpdateCategory
}

type CategoryResponse struct {
	CategoryId int    `json:"categoryId"`
	Title      string `json:"title"`
	ParentId   *int   `json:"parentId,omitempty"`
}

type GetCategoryRequest struct {
	CategoryId int
	AccountId  int
}

type GetCategoriesRequest struct {
	AccountId int
}

type DeleteCategoryRequest struct {
	CategoryId int
	AccountId  int
}
