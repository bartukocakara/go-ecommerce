package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindOrderItem() {
	c.OrderItemRepository = repository.NewOrderItemRepository(c.DB)
	c.OrderItemService = service.NewOrderItemService(c.OrderItemRepository)
	c.OrderItemHandler = handler.NewOrderItemHandler(c.OrderItemService)
}
