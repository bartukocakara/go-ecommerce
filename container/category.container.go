package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindCategory() {
	c.CategoryRepository = repository.NewCategoryRepository(c.DB)
	c.CategoryService = service.NewCategoryService(c.CategoryRepository)
	c.CategoryHandler = handler.NewCategoryHandler(c.CategoryService)
}
