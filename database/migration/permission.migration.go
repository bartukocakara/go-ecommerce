package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type PermissionMigration struct {
	DB *gorm.DB
}

func NewPermissionMigration(db *gorm.DB) *PermissionMigration {
	return &PermissionMigration{DB: db}
}

func (m *PermissionMigration) Migrate() {
	// Run the migration logic
	err := m.DB.AutoMigrate(&entity.Permission{})
	if err != nil {
		fmt.Println("Failed to migrate Permission table:", err)
		return
	}

	fmt.Println("Permission migration completed successfully")
}
