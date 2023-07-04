package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindRole() {
	c.RoleRepository = repository.NewRoleRepository(c.DB)
	c.RoleService = service.NewRoleService(c.RoleRepository)
	c.RoleHandler = handler.NewRoleHandler(c.RoleService)
}
