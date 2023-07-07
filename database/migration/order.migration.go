package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type OrderMigration struct {
	DB *gorm.DB
}

func NewOrderMigration(db *gorm.DB) *OrderMigration {
	return &OrderMigration{DB: db}
}

func (m *OrderMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.Order{})
	if err != nil {
		fmt.Println("Failed to migrate Order table:", err)
		return
	}

	fmt.Println("Order migration completed successfully.")
}
