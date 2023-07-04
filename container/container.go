package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	DB             *gorm.DB
	UserHandler    handler.UserHandler
	UserService    service.UserService
	UserRepository repository.UserRepository
}

func NewContainer(db *gorm.DB) *Container {
	return &Container{
		DB: db,
	}
}

func (c *Container) BindRepositories() {
	c.UserRepository = repository.NewUserRepository(c.DB)
}

func (c *Container) BindServices() {
	c.UserService = service.NewUserService(c.UserRepository)
}

func (c *Container) BindHandlers() {
	c.UserHandler = handler.NewUserHandler(c.UserService)
}
