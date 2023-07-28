package response

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"

	"github.com/gofiber/fiber/v2"
)

// RegistrationResponse represents the response structure for registration
type RegistrationResponse struct {
	User        *entity.User
	Role        *entity.Role
	AccessToken string
	TokenType   string
	ExpiresIn   int
}

// CreateRegistrationResponse generates the custom registration response
func CreateRegistrationResponse(ctx *fiber.Ctx, resp *dto.RegistrationResponse) error {
	response := CreateResponse(fiber.StatusCreated, "OK", fiber.Map{
		"user":         resp.User,
		"role":         resp.Role,
		"access_token": resp.AccessToken,
		"token_type":   "bearer",
		"expires_in":   3600,
	})

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
