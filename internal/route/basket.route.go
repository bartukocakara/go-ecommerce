package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupBasketRoutes(app *fiber.App, basketHandler handler.BasketHandler) {
	api := app.Group("/baskets")

	// Define your Basket routes
	api.Get("/", basketHandler.List)
	api.Get("/:id", basketHandler.Show)
	api.Post("/", basketHandler.Create)
	api.Put("/:id", basketHandler.Update)
	api.Delete("/:id", basketHandler.Delete)
}
