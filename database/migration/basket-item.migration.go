package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type BasketItemMigration struct {
	DB *gorm.DB
}

func NewBasketItemMigration(db *gorm.DB) *BasketItemMigration {
	return &BasketItemMigration{DB: db}
}

func (m *BasketItemMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.BasketItem{})
	if err != nil {
		fmt.Println("Failed to migrate BasketItem table:", err)
		return
	}

	fmt.Println("BasketItem migration completed successfully.")
}
