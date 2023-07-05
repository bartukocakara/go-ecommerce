package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(app *fiber.App, categoryHandler handler.CategoryHandler) {
	api := app.Group("/api/v1")

	// Define your category routes
	api.Get("/categories", categoryHandler.GetCategories)
	api.Get("/categories/:id", categoryHandler.GetCategoryByID)
	api.Post("/categories", categoryHandler.CreateCategory)
	api.Put("/categories/:id", categoryHandler.UpdateCategory)
	api.Delete("/categories/:id", categoryHandler.DeleteCategory)
}
