package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ProductCategorySeeder struct {
	DB *gorm.DB
}

func NewProductCategorySeeder(db *gorm.DB) *ProductCategorySeeder {
	return &ProductCategorySeeder{DB: db}
}

func (s *ProductCategorySeeder) Run() {
	productCategory := []entity.ProductCategory{
		{Name: "ProductCategory 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "ProductCategory 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// Add more ProductCategory data as needed
	}

	err := s.DB.Create(&productCategory).Error
	if err != nil {
		fmt.Println("Failed to seed product-category:", err)
		return
	}

	fmt.Println("ProductCategory seeding completed successfully.")
}
