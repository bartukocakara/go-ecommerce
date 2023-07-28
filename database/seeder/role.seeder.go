package seeder

import (
	"ecommerce/internal/entity"
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
	// Implement your role seeding logic here

	roles := []entity.Role{
		{Name: "Admin"},
		{Name: "User"},
		{Name: "Customer"},
	}

	for _, role := range roles {
		result := s.DB.Create(&role)
		if result.Error != nil {
			fmt.Println("Failed to create role:", result.Error)
			return
		}
	}
	fmt.Println("Role seeding completed successfully")
}
