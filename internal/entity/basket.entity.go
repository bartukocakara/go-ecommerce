package entity

import "time"

type Basket struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	UserID      uint         `gorm:"not null" json:"user_id"`
	User        User         `gorm:"foreignkey:UserID;constraint:onDelete:CASCADE" json:"-"`
	TotalPrice  float64      `gorm:"not null" json:"total_price"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	BasketItems []BasketItem `gorm:"foreignKey:BasketID" json:"basket_items"`
}
