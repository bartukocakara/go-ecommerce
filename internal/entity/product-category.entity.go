package entity

import "time"

type ProductCategory struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
	// ProductID  uint `gorm:"not null" json:"product_id"`
	// Product    Role `gorm:"foreignkey:ProductID;constraint:onDelete:CASCADE" json:"-"`
	// CategoryID uint `gorm:"not null" json:"category_id"`
	// Category   Role `gorm:"foreignkey:CategoryID;constraint:onDelete:CASCADE" json:"-"`
	// Add other fields as per your requirements
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
