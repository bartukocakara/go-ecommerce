package dto

type ProductCategoryDTO struct {
	ID        uint   `json:"id"`
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}

type CreateProductCategoryDTO struct {
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}

type FilterProductCategoryDTO struct {
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}

type UpdateProductCategoryDTO struct {
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}


