package seeder

import (
	"ecommerce/internal/entity"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
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
	var users []entity.User

	for i := 0; i < 100; i++ {
		firstName := faker.FirstName()
		lastName := faker.LastName()

		user := entity.User{
			FirstName: firstName,
			LastName:  lastName,
			Password:  "password",
			Email:     fmt.Sprintf("%s.%s@example.com", firstName, lastName),
			RoleID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		users = append(users, user)
	}

	for _, user := range users {
		result := s.DB.Create(&user)
		if result.Error != nil {
			fmt.Println("Failed to create user:", result.Error)
			return
		}
	}
	fmt.Printf("Successfully seeded %d users\n", 100)
}
