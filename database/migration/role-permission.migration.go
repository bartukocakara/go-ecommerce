package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type RolePermissionMigration struct {
	DB *gorm.DB
}

func NewRolePermissionMigration(db *gorm.DB) *RolePermissionMigration {
	return &RolePermissionMigration{DB: db}
}

func (m *RolePermissionMigration) Migrate() {
	// Run the migration logic
	err := m.DB.AutoMigrate(&entity.RolePermission{})
	if err != nil {
		fmt.Println("Failed to migrate Role Permission table:", err)
		return
	}

	fmt.Println("Role Permission migration completed successfully")
}
