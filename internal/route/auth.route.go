package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, handler handler.AuthHandler) {
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}
