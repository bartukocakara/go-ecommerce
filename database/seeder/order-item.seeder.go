package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrderItemSeeder struct {
	DB *gorm.DB
}

func NewOrderItemSeeder(db *gorm.DB) *OrderItemSeeder {
	return &OrderItemSeeder{DB: db}
}

func (s *OrderItemSeeder) Run() {
	orderItem := []entity.OrderItem{
		{Name: "OrderItem 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "OrderItem 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// Add more OrderItem data as needed
	}

	err := s.DB.Create(&orderItem).Error
	if err != nil {
		fmt.Println("Failed to seed order-item:", err)
		return
	}

	fmt.Println("OrderItem seeding completed successfully.")
}
