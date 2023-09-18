package entity

import "time"

type BasketItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignkey:ProductID;constraint:onDelete:CASCADE" json:"-"`
	BasketID  uint      `gorm:"not null" json:"basket_id"`
	Basket    Basket    `gorm:"foreignkey:BasketID;constraint:onDelete:CASCADE" json:"-"`
	Price     float64   `gorm:"not null" json:"price"`
	Quantity  string    `gorm:"quantity" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
