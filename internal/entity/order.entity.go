package entity

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey"`
	Code       string    `gorm:"unique"`
	BasketID   uint      `gorm:"not null" json:"basket_id"`
	Basket     Basket    `gorm:"foreignkey:BasketID;constraint:onDelete:CASCADE" json:"-"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
