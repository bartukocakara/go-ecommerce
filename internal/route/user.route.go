package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHandler handler.UserHandler) {
	userGroup := app.Group("/users")

	userGroup.Get("/", userHandler.GetUsers)
	userGroup.Get("/:id", userHandler.GetUserByID)
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Put("/:id", userHandler.UpdateUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)
}
