package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RolePermission(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the authenticated user's role from the context or wherever it is stored
		// Example: role := c.Locals("user_role").(string)

		// Check if the authenticated user's role is allowed
		// Change the logic according to your application's role-permission system
		allowed := false
		// Example: if role == "admin" || role == "superuser" {
		//     allowed = true
		// }

		// If the user's role is not allowed, return a forbidden response
		if !allowed {
			return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
		}

		// Continue to the next middleware or route handler
		return c.Next()
	}
}
