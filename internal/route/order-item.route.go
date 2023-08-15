package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupOrderItemRoutes(app *fiber.App, orderItemHandler handler.OrderItemHandler) {
	route := app.Group("/order-item")

	route.Get("/", orderItemHandler.GetOrderItems)
	route.Get("/:id", orderItemHandler.GetOrderItemByID)
	route.Post("/", orderItemHandler.CreateOrderItem)
	route.Put("/:id", orderItemHandler.UpdateOrderItem)
	route.Delete("/:id", orderItemHandler.DeleteOrderItem)
}
