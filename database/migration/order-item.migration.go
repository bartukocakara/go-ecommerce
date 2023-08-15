package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type OrderItemMigration struct {
	DB *gorm.DB
}

func NewOrderItemMigration(db *gorm.DB) *OrderItemMigration {
	return &OrderItemMigration{DB: db}
}

func (m *OrderItemMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.OrderItem{})
	if err != nil {
		fmt.Println("Failed to migrate OrderItem table:", err)
		return
	}

	fmt.Println("OrderItem migration completed successfully.")
}
