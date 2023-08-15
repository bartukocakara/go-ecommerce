package entity

import "time"

type OrderItem struct {
	ID uint `gorm:"primaryKey"`
	// OrderID   uint `gorm:"not null" json:"order_id"`
	// Order     Role `gorm:"foreignkey:OrderID;constraint:onDelete:CASCADE" json:"-"`
	// ProductID uint `gorm:"not null" json:"product_id"`
	// Product   Role `gorm:"foreignkey:ProductID;constraint:onDelete:CASCADE" json:"-"`
	Name string `gorm:"unique"`
	// Add other fields as per your requirements
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
