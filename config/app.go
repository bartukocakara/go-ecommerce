package config

import (
	"ecommerce/container"
	"ecommerce/database/migration"
	"ecommerce/database/seeder"
	"ecommerce/internal/middleware"
	"ecommerce/internal/route"
	"ecommerce/internal/validation"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type stubTemplates struct {
	Handler    string
	Repository string
	Service    string
	Route      string
	Dto        string
	Entity     string
	Container  string
	Seeder     string
	Migration  string
}

const (
	handlerStubPath    = "./stubs/module.handler.stub"
	repositoryStubPath = "./stubs/module.repository.stub"
	serviceStubPath    = "./stubs/module.service.stub"
	routeStubPath      = "./stubs/module.route.stub"
	dtoStubPath        = "./stubs/module.dto.stub"
	entityStubPath     = "./stubs/module.entity.stub"
	containerStubPath  = "./stubs/module.container.stub"
	seederStubPath     = "./stubs/module.seeder.stub"
	migrationStubPath  = "./stubs/module.migration.stub"
)

func NewApp() (*fiber.App, error) {
	app := fiber.New()
	validate := NewValidator()
	if err := validation.LoadMessages(); err != nil {
		// Handle the error if loading messages fails
		panic(err)
	}
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("validator", validate)
		return c.Next()
	})
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
	c := container.NewContainer(db)
	c.BindAuth()

	route.SetupAuthRoutes(app, c.AuthHandler)

	app.Use(middleware.JWT())

	// Bind dependencies in the container
	c.BindUser()
	c.BindRole()
	c.BindForgotPasswordToken()
	c.BindProduct()
	c.BindCategory()
	c.BindBasket()
	c.BindBasketItem()
	c.BindOrder()

	// Setup routes
	route.SetupUserRoutes(app, c.UserHandler, db)
	route.SetupRoleRoutes(app, c.RoleHandler)
	route.SetupProductRoutes(app, c.ProductHandler)
	route.SetupCategoryRoutes(app, c.CategoryHandler)
	route.SetupBasketRoutes(app, c.BasketHandler)
	route.SetupBasketItemRoutes(app, c.BasketItemHandler)
	route.SetupOrderRoutes(app, c.OrderHandler)

	commands(db)

	app.Shutdown()

	return app, nil
}

func commands(db *gorm.DB) {
	args := os.Args
	if len(args) > 1 {
		command := strings.ToLower(args[1])

		switch command {
		case "seed":
			runSeeders(db)
		case "migrate":
			runMigrations(db)
		case "generate-module":
			if len(args) > 2 {
				moduleName := args[2]
				templates, err := loadStubTemplates()
				if err != nil {
					fmt.Println("Failed to get stub templates:", err)
					return
				}
				generateFiles(moduleName, templates)
			} else {
				fmt.Println("Missing module name. Usage: go run main.go generate-file [module]")
			}
		case "kill":
			fmt.Println("Shutting down the app...")
			// Perform any cleanup or graceful shutdown operations here if needed.
			// Then exit the application.
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Usage: go run main.go [seed|migrate|generate-file]")
		}
	}
}


func runSeeders(db *gorm.DB) {
	// Run your seeders here
	// Example:
	userSeeder := seeder.NewUserSeeder(db)
	roleSeeder := seeder.NewRoleSeeder(db)
	permissionSeeder := seeder.NewPermissionSeeder(db)
	rolePermissionSeeder := seeder.NewRolePermissionSeeder(db)
	productSeeder := seeder.NewProductSeeder(db)
	categorySeeder := seeder.NewCategorySeeder(db)
	userSeeder.Run()
	roleSeeder.Run()
	permissionSeeder.Run()
	rolePermissionSeeder.Run()
	productSeeder.Run()
	categorySeeder.Run()
	fmt.Println("Running seeders...")
}

func runMigrations(db *gorm.DB) {
	// Run your migrations here
	// Example:
	userMigration := migration.NewUserMigration(db)
	roleMigration := migration.NewRoleMigration(db)
	permissionMigration := migration.NewPermissionMigration(db)
	rolePermissionMigration := migration.NewRolePermissionMigration(db)
	productMigration := migration.NewProductMigration(db)
	categoryMigration := migration.NewCategoryMigration(db)
	forgotPasswordTokenMigration := migration.NewForgotPasswordTokenMigration(db)
	userMigration.Migrate()
	roleMigration.Migrate()
	permissionMigration.Migrate()
	rolePermissionMigration.Migrate()
	productMigration.Migrate()
	categoryMigration.Migrate()
	forgotPasswordTokenMigration.Migrate()

	fmt.Println("Running migrations...")
}


func NewValidator() *validator.Validate {
	return validator.New()
}
