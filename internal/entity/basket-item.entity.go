package entity

import "time"

type BasketItem struct {
	ID        uint      `gorm:"primaryKey"`
	Quantity  string    `gorm:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
