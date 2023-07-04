package seeder

import (
	"fmt"

	"gorm.io/gorm"
)

type RoleSeeder struct {
	DB *gorm.DB
}

func NewRoleSeeder(db *gorm.DB) *RoleSeeder {
	return &RoleSeeder{DB: db}
}

func (s *RoleSeeder) Run() {
	// Implement your role seeder logic here
	fmt.Println("Running role seeder...")
}
