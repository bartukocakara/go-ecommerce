package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ProductSeeder struct {
	DB *gorm.DB
}

func NewProductSeeder(db *gorm.DB) *ProductSeeder {
	return &ProductSeeder{DB: db}
}

func (s *ProductSeeder) Run() {
	products := []entity.Product{
		{Name: "Product 1", Price: 10.99, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Product 2", Price: 19.99, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// Add more product data as needed
	}

	err := s.DB.Create(&products).Error
	if err != nil {
		fmt.Println("Failed to seed products:", err)
		return
	}

	fmt.Println("Product seeding completed successfully.")
}
