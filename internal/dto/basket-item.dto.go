package dto

type CreateBasketItemDto struct {
	Quantity string `json:"quantity"`
}

type UpdateBasketItemDto struct {
	Quantity string `json:"quantity"`
}
