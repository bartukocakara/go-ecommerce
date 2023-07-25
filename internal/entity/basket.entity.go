package entity

import "time"

type Basket struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	TotalPrice float64   `gorm:"unique;not null" json:"total_price"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
