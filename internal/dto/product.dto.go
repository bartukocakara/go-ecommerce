package dto

type CreateProductDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,min=1"`
}

type UpdateProductDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type FilterProductDto struct {
	Name string `json:"name"`
}
