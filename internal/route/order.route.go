package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(app *fiber.App, orderHandler handler.OrderHandler) {
	route := app.Group("/order")

	route.Get("/", orderHandler.GetOrders)
	route.Get("/:id", orderHandler.GetOrderByID)
	route.Post("/", orderHandler.CreateOrder)
	route.Put("/:id", orderHandler.UpdateOrder)
	route.Delete("/:id", orderHandler.DeleteOrder)
}
