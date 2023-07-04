package entity

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `gorm:"unique;not null"`
	LastName  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}