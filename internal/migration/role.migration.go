package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type RoleMigration struct {
	DB *gorm.DB
}

func NewRoleMigration(db *gorm.DB) *RoleMigration {
	return &RoleMigration{DB: db}
}

func (m *RoleMigration) Migrate() {
	// Run the migration logic
	err := m.DB.AutoMigrate(&entity.Role{})
	if err != nil {
		fmt.Println("Failed to migrate Role table:", err)
		return
	}

	fmt.Println("Role migration completed successfully")
}
