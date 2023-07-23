package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, handler handler.AuthHandler) {
	authGroup := app.Group("/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
	authGroup.Post("/auth/forget-password", handler.ForgetPassword)
}
