package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type CategorySeeder struct {
	DB *gorm.DB
}

func NewCategorySeeder(db *gorm.DB) *CategorySeeder {
	return &CategorySeeder{DB: db}
}

func (s *CategorySeeder) Run() {
	var categories []entity.Category

	for i := 0; i < 10; i++ {
		category := entity.Category{
			Name:      faker.Word(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		categories = append(categories, category)
	}

	err := s.DB.Create(&categories).Error
	if err != nil {
		fmt.Println("Failed to seed categories:", err)
		return
	}

	fmt.Println("Category seeding completed successfully.")
}
