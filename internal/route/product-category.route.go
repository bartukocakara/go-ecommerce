package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupProductCategoryRoutes(app *fiber.App, productCategoryHandler handler.ProductCategoryHandler) {
	route := app.Group("/product-category")

	route.Get("/", productCategoryHandler.GetProductCategorys)
	route.Get("/:id", productCategoryHandler.GetProductCategoryByID)
	route.Post("/", productCategoryHandler.CreateProductCategory)
	route.Put("/:id", productCategoryHandler.UpdateProductCategory)
	route.Delete("/:id", productCategoryHandler.DeleteProductCategory)
}
