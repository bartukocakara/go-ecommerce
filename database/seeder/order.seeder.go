package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrderSeeder struct {
	DB *gorm.DB
}

func NewOrderSeeder(db *gorm.DB) *OrderSeeder {
	return &OrderSeeder{DB: db}
}

func (s *OrderSeeder) Run() {
	order := []entity.Order{
		{Name: "Order 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Order 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// Add more Order data as needed
	}

	err := s.DB.Create(&order).Error
	if err != nil {
		fmt.Println("Failed to seed order:", err)
		return
	}

	fmt.Println("Order seeding completed successfully.")
}
