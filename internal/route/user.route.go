package route

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHandler handler.UserHandler) {
	userGroup := app.Group("/users")

	userGroup.Get("/", middleware.RolePermission([]string{"Admin", "User"}), userHandler.GetUsers)
	userGroup.Get("/:id", userHandler.GetUserByID)
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Put("/:id", userHandler.UpdateUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)
}
