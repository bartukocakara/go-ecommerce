package main

import (
	"embed"
	"log"

	"ecommerce/config"
	"ecommerce/internal/middleware"
)

//go:embed stubs/*.stub
var stubs embed.FS

func main() {
	// Create a new Fiber app instance
	app, err := config.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	// Apply middlewares
	app.Use(middleware.CORS())
	app.Use(middleware.JWT())
	app.Use(middleware.RolePermission([]string{"admin", "superuser"}))

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
