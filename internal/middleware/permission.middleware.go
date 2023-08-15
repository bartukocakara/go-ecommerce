package middleware

import (
	"ecommerce/internal/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PermissionMiddleware(allowedPermission string, userRepo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user's role from the context
		// userRole := c.Locals("userRole").(string)
		userID := c.Locals("user_id").(uint)
		fmt.Print(userID)
		userPermissions, err := userRepo.GetPermissionsByUserID(userID)
		if err != nil {
			return createErrorResponse(c, fiber.StatusInternalServerError, "Failed to get user role")
		}
		// Check if the user's role is allowed
		if !contains(userPermissions, allowedPermission) {
			return createErrorResponse(c, fiber.StatusForbidden, "Permission denied")
		}

		return c.Next()
	}
}
