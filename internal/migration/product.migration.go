package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type ProductMigration struct {
	DB *gorm.DB
}

func NewProductMigration(db *gorm.DB) *ProductMigration {
	return &ProductMigration{DB: db}
}

func (m *ProductMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.Product{})
	if err != nil {
		fmt.Println("Failed to migrate product table:", err)
		return
	}

	fmt.Println("Product migration completed successfully.")
}
