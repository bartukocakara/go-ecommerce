package main

import (
	"log"

	"ecommerce/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(middleware.CORS())

	log.Fatal(app.Listen(":3000"))
}
