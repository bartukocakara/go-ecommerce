package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CategorySeeder struct {
	DB *gorm.DB
}

func NewCategorySeeder(db *gorm.DB) *CategorySeeder {
	return &CategorySeeder{DB: db}
}

func (s *CategorySeeder) Run() {
	categories := []entity.Category{
		{Name: "Category 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Category 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		// Add more category data as needed
	}

	err := s.DB.Create(&categories).Error
	if err != nil {
		fmt.Println("Failed to seed categories:", err)
		return
	}

	fmt.Println("Category seeding completed successfully.")
}
