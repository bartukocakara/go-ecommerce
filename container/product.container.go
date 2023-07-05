package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindProduct() {
	c.ProductRepository = repository.NewProductRepository(c.DB)
	c.ProductService = service.NewProductService(c.ProductRepository)
	c.ProductHandler = handler.NewProductHandler(c.ProductService)
}
