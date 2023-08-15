package dto

type OrderItemDTO struct {
	ID        uint   `json:"id"`
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}

type CreateOrderItemDTO struct {
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}

type FilterOrderItemDTO struct {
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}

type UpdateOrderItemDTO struct {
	Name string `json:"name"`
	// Add other filter fields as per your requirements
}


