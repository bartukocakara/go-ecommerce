package route

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/middleware"
	"ecommerce/internal/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(app *fiber.App, userHandler handler.UserHandler, db *gorm.DB) {
	userGroup := app.Group("/users")

	userGroup.Get("/", middleware.RoleMiddleware([]string{"Admin", "User"},
		repository.NewUserRepository(db)),
		// middleware.PermissionMiddleware("list_user",
		// 	repository.NewUserRepository(db)),
		userHandler.List)
	userGroup.Get("/:id", userHandler.Show)
	userGroup.Post("/", userHandler.Create)
	userGroup.Put("/:id", userHandler.Update)
	userGroup.Delete("/:id", userHandler.Delete)
}
