package seeder

import (
	"fmt"

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
	fmt.Println("Running user seeder...")
}
