package config

import (
	"ecommerce/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewApp() (*fiber.App, error) {
	app := fiber.New()

	// Connect to the database
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil, err
	}

	// Set up the database connection in the context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Middleware setup
	app.Use(middleware.CORS())
	app.Use(middleware.JWT())
	app.Use(middleware.RolePermission([]string{"admin", "superuser"}))

	// Rest of the code...

	// Close the database connection when the application exits
	app.Shutdown()

	return app, nil
}
