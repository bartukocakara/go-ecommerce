package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindUser() {
	c.UserRepository = repository.NewUserRepository(c.DB)
	c.UserService = service.NewUserService(c.UserRepository)
	c.UserHandler = handler.NewUserHandler(c.UserService)
}
