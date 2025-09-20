package request

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}
