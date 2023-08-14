package route

import (
	"ecommerce/internal/handler"
	validator "ecommerce/internal/middleware/validation"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, handler handler.AuthHandler) {
	authGroup := app.Group("/auth")
	authGroup.Post("/register", validator.RegisterValidator, handler.Register)
	authGroup.Post("/login", handler.Login)
	authGroup.Post("/forgot-password", handler.ForgotPassword)
}
