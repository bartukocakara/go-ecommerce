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
	RoleHandler    handler.RoleHandler
	RoleService    service.RoleService
	RoleRepository repository.RoleRepository
	AuthHandler    handler.AuthHandler
	AuthService    service.AuthService
}

func NewContainer(db *gorm.DB) *Container {
	return &Container{
		DB: db,
	}
}

func (c *Container) BindRepositories() {
	c.UserRepository = repository.NewUserRepository(c.DB)
	c.RoleRepository = repository.NewRoleRepository(c.DB)
}

func (c *Container) BindServices() {
	c.UserService = service.NewUserService(c.UserRepository)
	c.RoleService = service.NewRoleService(c.RoleRepository)
}

func (c *Container) BindHandlers() {
	c.UserHandler = handler.NewUserHandler(c.UserService)
	c.RoleHandler = handler.NewRoleHandler(c.RoleService)
}
