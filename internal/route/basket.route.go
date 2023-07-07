package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupBasketRoutes(app *fiber.App, basketHandler handler.BasketHandler) {
	api := app.Group("/api/v1")

	// Define your Basket routes
	api.Get("/baskets", basketHandler.GetBaskets)
	api.Get("/baskets/:id", basketHandler.GetBasketByID)
	api.Post("/baskets", basketHandler.CreateBasket)
	api.Put("/baskets/:id", basketHandler.UpdateBasket)
	api.Delete("/baskets/:id", basketHandler.DeleteBasket)
}
