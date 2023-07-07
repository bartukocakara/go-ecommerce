package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type CategoryMigration struct {
	DB *gorm.DB
}

func NewCategoryMigration(db *gorm.DB) *CategoryMigration {
	return &CategoryMigration{DB: db}
}

func (m *CategoryMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.Category{})
	if err != nil {
		fmt.Println("Failed to migrate category table:", err)
		return
	}

	fmt.Println("Category migration completed successfully.")
}
