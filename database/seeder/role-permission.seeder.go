package seeder

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type RolePermissionSeeder struct {
	DB *gorm.DB
}

func NewRolePermissionSeeder(db *gorm.DB) *RolePermissionSeeder {
	return &RolePermissionSeeder{DB: db}
}

func (s *RolePermissionSeeder) Run() {
	// Implement your RolePermission seeding logic here

	rolePermissions := []entity.RolePermission{
		{RoleID: 1, PermissionID: 1},
		{RoleID: 1, PermissionID: 2},
		{RoleID: 1, PermissionID: 3},
		{RoleID: 1, PermissionID: 4},
		{RoleID: 1, PermissionID: 5},
		{RoleID: 1, PermissionID: 6},
		{RoleID: 1, PermissionID: 7},
		{RoleID: 1, PermissionID: 8},
		{RoleID: 1, PermissionID: 9},
		{RoleID: 1, PermissionID: 10},
		{RoleID: 1, PermissionID: 11},

		{RoleID: 2, PermissionID: 1},
		{RoleID: 2, PermissionID: 2},
		{RoleID: 2, PermissionID: 3},
		{RoleID: 2, PermissionID: 4},
		{RoleID: 2, PermissionID: 5},
		{RoleID: 2, PermissionID: 6},
		{RoleID: 2, PermissionID: 7},
		{RoleID: 2, PermissionID: 8},
		{RoleID: 2, PermissionID: 9},
	}

	for _, rolePermission := range rolePermissions {
		result := s.DB.Create(&rolePermission)
		if result.Error != nil {
			fmt.Println("Failed to create role permission:", result.Error)
			return
		}
	}
	fmt.Println("Role Permission seeding completed successfully")
}
