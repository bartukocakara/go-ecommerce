package dto

type CreateCategoryDto struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryDto struct {
	Name string `json:"name"`
}
