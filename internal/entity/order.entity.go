package entity

import "time"

type Order struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique"`
	// Add other fields as per your requirements
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
