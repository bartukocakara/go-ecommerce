package main

import (
	"embed"
	"log"

	"ecommerce/config"
)

//go:embed stubs/*.stub
var stubs embed.FS

func main() {
	// Create a new Fiber app instance
	app, err := config.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	// Start the server
	log.Fatal(app.Listen(":8000"))
}
