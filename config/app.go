package config

import (
	"ecommerce/container"
	"ecommerce/internal/middleware"
	"ecommerce/internal/migration"
	"ecommerce/internal/route"
	"ecommerce/internal/seeder"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	// Create a new container instance with the database connection
	c := container.NewContainer(db)

	// Bind dependencies in the container
	c.BindUser()
	c.BindRole()

	// Setup routes
	route.SetupUserRoutes(app, c.UserHandler)
	route.SetupRoleRoutes(app, c.RoleHandler)

	runSeedAndMigrationCommands(db)

	// Rest of the code...

	// Close the database connection when the application exits
	app.Shutdown()

	return app, nil
}

func runSeedAndMigrationCommands(db *gorm.DB) {
	args := os.Args
	if len(args) > 1 {
		command := strings.ToLower(args[1])

		switch command {
		case "seed":
			runSeeders(db)
		case "migrate":
			runMigrations(db)
		default:
			fmt.Println("Invalid command. Usage: go run main.go [seed|migrate]")
		}
	}
}

func runSeeders(db *gorm.DB) {
	// Run your seeders here
	// Example:
	userSeeder := seeder.NewUserSeeder(db)
	userSeeder.Run()

	fmt.Println("Running seeders...")
}

func runMigrations(db *gorm.DB) {
	// Run your migrations here
	// Example:
	userMigration := migration.NewUserMigration(db)
	userMigration.Migrate()

	fmt.Println("Running migrations...")
}
