package middleware

import (
	"ecommerce/internal/handler/response"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RolePermission(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user's role from the context
		userRole := c.Locals("userRole").(string)
		fmt.Print(userRole)
		// Check if the user's role is allowed
		if !contains(allowedRoles, userRole) {
			return createErrorResponse(c, fiber.StatusForbidden, "Permission denied")
		}

		return c.Next()
	}
}

func contains(slice []string, search string) bool {
	for _, value := range slice {
		if strings.ToLower(value) == strings.ToLower(search) {
			return true
		}
	}
	return false
}

func createErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.JSON(response.GenericErrorResponse{
		Message:    message,
		StatusCode: status,
	})
}
