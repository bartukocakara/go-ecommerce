package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindAuth() {
	c.UserRepository = repository.NewUserRepository(c.DB)
	c.ForgotPasswordRepository = repository.NewForgotPasswordTokenRepository(c.DB)
	c.AuthService = *service.NewAuthService(c.UserRepository, c.ForgotPasswordRepository)
	c.AuthHandler = *handler.NewAuthHandler(c.AuthService)
}
