package entity

import (
	"time"
)

type Product struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
