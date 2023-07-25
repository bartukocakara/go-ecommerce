package dto

type CreateBasketDto struct {
	TotalPrice float64 `json:"total_price"`
}

type UpdateBasketDto struct {
	TotalPrice float64 `json:"total_price"`
}
