package route

import (
	"ecommerce/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(app *fiber.App, roleHandler handler.RoleHandler) {
	roleGroup := app.Group("/roles")

	roleGroup.Get("/", roleHandler.GetRoles)
	roleGroup.Get("/:id", roleHandler.GetRoleByID)
	roleGroup.Post("/", roleHandler.CreateRole)
	roleGroup.Put("/:id", roleHandler.UpdateRole)
	roleGroup.Delete("/:id", roleHandler.DeleteRole)
}
