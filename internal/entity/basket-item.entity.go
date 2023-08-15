package entity

import "time"

type BasketItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	BasketID  uint      `gorm:"not null" json:"basket_id"`
	Basket    Role      `gorm:"foreignkey:BasketID;constraint:onDelete:CASCADE" json:"-"`
	Quantity  string    `gorm:"quantity" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
