package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type ProductCategoryMigration struct {
	DB *gorm.DB
}

func NewProductCategoryMigration(db *gorm.DB) *ProductCategoryMigration {
	return &ProductCategoryMigration{DB: db}
}

func (m *ProductCategoryMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.ProductCategory{})
	if err != nil {
		fmt.Println("Failed to migrate ProductCategory table:", err)
		return
	}

	fmt.Println("ProductCategory migration completed successfully.")
}
