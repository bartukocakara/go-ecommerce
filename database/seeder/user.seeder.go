package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
}

func (s *UserSeeder) Run() {
	// Implement your user seeder logic here
	users := []entity.User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Password:  "password1",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			FirstName: "Jane",
			LastName:  "Smith",
			Password:  "password2",
			Email:     "jane@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		// Add more users as needed
	}

	for _, user := range users {
		result := s.DB.Create(&user)
		if result.Error != nil {
			fmt.Println("Failed to create user:", result.Error)
			return
		}
	}
	fmt.Println("Running user seeder...")
}
