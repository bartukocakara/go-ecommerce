package dto

type CreateProductDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type UpdateProductDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type FilterProductDto struct {
	Name string `json:"name"`
}
