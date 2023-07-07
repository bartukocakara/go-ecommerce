package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindBasketItem() {
	c.BasketItemRepository = repository.NewBasketItemRepository(c.DB)
	c.BasketItemService = service.NewBasketItemService(c.BasketItemRepository)
	c.BasketItemHandler = handler.NewBasketItemHandler(c.BasketItemService)
}
