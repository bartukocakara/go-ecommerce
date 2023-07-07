package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
)

func (c *Container) BindOrder() {
	c.OrderRepository = repository.NewOrderRepository(c.DB)
	c.OrderService = service.NewOrderService(c.OrderRepository)
	c.OrderHandler = handler.NewOrderHandler(c.OrderService)
}
