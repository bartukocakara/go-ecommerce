package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindProductCategory() {
	c.ProductCategoryRepository = repository.NewProductCategoryRepository(c.DB)
	c.ProductCategoryService = service.NewProductCategoryService(c.ProductCategoryRepository)
	c.ProductCategoryHandler = handler.NewProductCategoryHandler(c.ProductCategoryService)
}
