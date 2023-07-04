package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserMigration struct {
	DB *gorm.DB
}

func NewUserMigration(db *gorm.DB) *UserMigration {
	return &UserMigration{DB: db}
}

func (m *UserMigration) Migrate() {
	// Run the migration logic
	err := m.DB.AutoMigrate(&entity.User{})
	if err != nil {
		fmt.Println("Failed to migrate User table:", err)
		return
	}

	fmt.Println("User migration completed successfully")
}
