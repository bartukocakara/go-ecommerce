package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"

	"gorm.io/gorm"
)

type AuthContainer struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
	AuthService    service.AuthService
	AuthHandler    handler.AuthHandler
}

func NewAuthContainer(db *gorm.DB) *AuthContainer {
	return &AuthContainer{
		DB: db,
	}
}

func (c *Container) BindAuth() {
	c.UserRepository = repository.NewUserRepository(c.DB)
	c.UserService = service.NewUserService(c.UserRepository)
	c.UserHandler = handler.NewUserHandler(c.UserService)
}
