package config

import (
	"ecommerce/container"
	"ecommerce/database/migration"
	"ecommerce/database/seeder"
	"ecommerce/internal/middleware"
	"ecommerce/internal/route"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

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
	// app.Use(middleware.RolePermission([]string{"admin", "superuser"}))

	// Create a new container instance with the database connection

	// Bind dependencies in the container
	c.BindUser()
	c.BindRole()
	c.BindProduct()
	c.BindCategory()
	c.BindBasket()
	c.BindBasketItem()
	c.BindOrder()

	// Setup routes
	route.SetupUserRoutes(app, c.UserHandler)
	route.SetupRoleRoutes(app, c.RoleHandler)
	route.SetupProductRoutes(app, c.ProductHandler)
	route.SetupCategoryRoutes(app, c.CategoryHandler)
	route.SetupBasketRoutes(app, c.BasketHandler)
	route.SetupBasketItemRoutes(app, c.BasketItemHandler)
	route.SetupOrderRoutes(app, c.OrderHandler)

	commands(db)

	// Rest of the code...

	// Close the database connection when the application exits
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
		case "generate-file":
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
		default:
			fmt.Println("Invalid command. Usage: go run main.go [seed|migrate|generate-file]")
		}
	}
}

type StubData struct {
	ModuleVar   string
	ModuleTitle string
}

func loadTemplateContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func loadStubTemplates() (*stubTemplates, error) {
	templates := &stubTemplates{}

	repositoryContent, err := loadTemplateContent(repositoryStubPath)
	if err != nil {
		return nil, err
	}
	templates.Repository = repositoryContent

	serviceContent, err := loadTemplateContent(serviceStubPath)
	if err != nil {
		return nil, err
	}
	templates.Service = serviceContent

	routeContent, err := loadTemplateContent(routeStubPath)
	if err != nil {
		return nil, err
	}
	templates.Route = routeContent

	dtoContent, err := loadTemplateContent(dtoStubPath)
	if err != nil {
		return nil, err
	}
	templates.Dto = dtoContent

	entityContent, err := loadTemplateContent(entityStubPath)
	if err != nil {
		return nil, err
	}
	templates.Entity = entityContent

	containerContent, err := loadTemplateContent(containerStubPath)
	if err != nil {
		return nil, err
	}
	templates.Container = containerContent

	handlerContent, err := loadTemplateContent(handlerStubPath)
	if err != nil {
		return nil, err
	}
	templates.Handler = handlerContent

	migrationContent, err := loadTemplateContent(migrationStubPath)
	if err != nil {
		return nil, err
	}
	templates.Migration = migrationContent

	seederContent, err := loadTemplateContent(seederStubPath)
	if err != nil {
		return nil, err
	}
	templates.Seeder = seederContent

	return templates, nil
}

func runSeeders(db *gorm.DB) {
	// Run your seeders here
	// Example:
	userSeeder := seeder.NewUserSeeder(db)
	roleSeeder := seeder.NewRoleSeeder(db)
	productSeeder := seeder.NewProductSeeder(db)
	categorySeeder := seeder.NewCategorySeeder(db)
	userSeeder.Run()
	roleSeeder.Run()
	productSeeder.Run()
	categorySeeder.Run()
	fmt.Println("Running seeders...")
}

func runMigrations(db *gorm.DB) {
	// Run your migrations here
	// Example:
	userMigration := migration.NewUserMigration(db)
	roleMigration := migration.NewRoleMigration(db)
	productMigration := migration.NewProductMigration(db)
	categoryMigration := migration.NewCategoryMigration(db)
	userMigration.Migrate()
	roleMigration.Migrate()
	productMigration.Migrate()
	categoryMigration.Migrate()

	fmt.Println("Running migrations...")
}

func createFile(filePath, content string) error {
	// Create the directory path if it doesn't exist
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// Generate files for the specified module using the provided stub templates
func generateFiles(moduleName string, templates *stubTemplates) {
	fmt.Println("Generating files for", moduleName)

	// Convert the module name to title case
	moduleTitle := strings.Title(moduleName)

	// Define the file names and content for the module
	files := map[string]string{
		fmt.Sprintf("internal/handler/%s.handler.go", moduleName):       generateContent(templates.Handler, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("internal/repository/%s.repository.go", moduleName): generateContent(templates.Repository, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("internal/service/%s.service.go", moduleName):       generateContent(templates.Service, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("internal/route/%s.route.go", moduleName):           generateContent(templates.Route, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("internal/entity/%s.entity.go", moduleName):         generateContent(templates.Entity, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("internal/dto/%s.dto.go", moduleName):               generateContent(templates.Dto, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("container/%s.container.go", moduleName):            generateContent(templates.Container, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("database/seeder/%s.seeder.go", moduleName):         generateContent(templates.Seeder, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
		fmt.Sprintf("database/migration/%s.migration.go", moduleName):   generateContent(templates.Migration, StubData{ModuleVar: moduleName, ModuleTitle: moduleTitle}),
	}

	// Generate the files
	for filePath, content := range files {
		err := createFile(filePath, content)
		if err != nil {
			fmt.Println("Failed to generate file:", err)
		} else {
			fmt.Println("Generated file:", filePath)
		}
	}
}

func generateContent(templateContent string, data StubData) string {
	tmpl, err := template.New("stub").Parse(templateContent)
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return ""
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return ""
	}

	return buf.String()
}
