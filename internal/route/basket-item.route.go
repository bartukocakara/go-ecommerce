package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupBasketItemRoutes(app *fiber.App, basketItemHandler handler.BasketItemHandler) {
	api := app.Group("/api/v1")

	// Define your Basket routes
	api.Get("/basket-items", basketItemHandler.List)
	api.Get("/basket-items/:id", basketItemHandler.Show)
	api.Post("/basket-items", basketItemHandler.Create)
	api.Put("/basket-items/:id", basketItemHandler.Update)
	api.Delete("/basket-items/:id", basketItemHandler.Delete)
}
