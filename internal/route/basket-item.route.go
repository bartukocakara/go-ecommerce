package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupBasketItemRoutes(app *fiber.App, basketItemHandler handler.BasketItemHandler) {
	api := app.Group("/api/v1")

	// Define your Basket routes
	api.Get("/basket-items", basketItemHandler.GetBasketItems)
	api.Get("/basket-items/:id", basketItemHandler.GetBasketItemByID)
	api.Post("/basket-items", basketItemHandler.CreateBasketItem)
	api.Put("/basket-items/:id", basketItemHandler.UpdateBasketItem)
	api.Delete("/basket-items/:id", basketItemHandler.DeleteBasketItem)
}
