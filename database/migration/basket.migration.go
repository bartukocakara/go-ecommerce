package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type BasketMigration struct {
	DB *gorm.DB
}

func NewBasketMigration(db *gorm.DB) *BasketMigration {
	return &BasketMigration{DB: db}
}

func (m *BasketMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.Basket{})
	if err != nil {
		fmt.Println("Failed to migrate Basket table:", err)
		return
	}

	fmt.Println("Basket migration completed successfully.")
}
