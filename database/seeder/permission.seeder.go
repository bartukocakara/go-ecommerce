package seeder

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type PermissionSeeder struct {
	DB *gorm.DB
}

func NewPermissionSeeder(db *gorm.DB) *PermissionSeeder {
	return &PermissionSeeder{DB: db}
}

func (s *PermissionSeeder) Run() {
	// Implement your permission seeding logic here

	permissions := []entity.Permission{
		{Name: "list_product"},
		{Name: "create_product"},
		{Name: "show_product"},
		{Name: "update_product"},
		{Name: "delete_product"},

		{Name: "list_category"},
		{Name: "create_category"},
		{Name: "show_category"},
		{Name: "update_category"},
		{Name: "delete_category"},

		{Name: "list_basket"},
		{Name: "create_basket"},
		{Name: "show_basket"},
		{Name: "update_basket"},
		{Name: "delete_basket"},

		{Name: "list_order"},
		{Name: "create_order"},
		{Name: "show_order"},
		{Name: "update_order"},
		{Name: "delete_order"},
	}

	for _, permission := range permissions {
		result := s.DB.Create(&permission)
		if result.Error != nil {
			fmt.Println("Failed to create permission:", result.Error)
			return
		}
	}
	fmt.Println("Permission seeding completed successfully")
}
