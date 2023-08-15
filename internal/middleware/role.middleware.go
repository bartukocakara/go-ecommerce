package middleware

import (
	"ecommerce/internal/handler/response"
	"ecommerce/internal/repository"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(allowedRoles []string, userRepo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user's role from the context
		// userRole := c.Locals("userRole").(string)
		userID := c.Locals("user_id").(uint)
		fmt.Print(userID)
		userRolName, err := userRepo.GetUserRoleNameByID(userID)
		if err != nil {
			return createErrorResponse(c, fiber.StatusInternalServerError, "Failed to get user role")
		}
		// Check if the user's role is allowed
		if !contains(allowedRoles, userRolName) {
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
