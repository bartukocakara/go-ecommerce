package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindBasket() {
	c.BasketRepository = repository.NewBasketRepository(c.DB)
	c.BasketService = service.NewBasketService(c.BasketRepository)
	c.BasketHandler = handler.NewBasketHandler(c.BasketService)
}
