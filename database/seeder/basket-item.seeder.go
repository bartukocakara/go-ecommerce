package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BasketItemSeeder struct {
	DB *gorm.DB
}

func NewBasketItemSeeder(db *gorm.DB) *BasketItemSeeder {
	return &BasketItemSeeder{DB: db}
}

func (s *BasketItemSeeder) Run() {
	basketItems := []entity.BasketItem{
		{Quantity: "BasketItem 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Quantity: "BasketItem 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// Add more BasketItem data as needed
	}

	err := s.DB.Create(&basketItems).Error
	if err != nil {
		fmt.Println("Failed to seed basketItems:", err)
		return
	}

	fmt.Println("BasketItem seeding completed successfully.")
}
