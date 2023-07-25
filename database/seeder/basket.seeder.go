package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BasketSeeder struct {
	DB *gorm.DB
}

func NewBasketSeeder(db *gorm.DB) *BasketSeeder {
	return &BasketSeeder{DB: db}
}

func (s *BasketSeeder) Run() {
	baskets := []entity.Basket{
		{TotalPrice: 19.50, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{TotalPrice: 19.50, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// Add more Basket data as needed
	}

	err := s.DB.Create(&baskets).Error
	if err != nil {
		fmt.Println("Failed to seed baskets:", err)
		return
	}

	fmt.Println("Basket seeding completed successfully.")
}
