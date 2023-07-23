package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindAuth() {
	c.UserRepository = repository.NewUserRepository(c.DB)
	c.AuthService = *service.NewAuthService(c.UserRepository)
	c.AuthHandler = *handler.NewAuthHandler(c.AuthService)
}