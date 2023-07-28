package validator

import (
	"ecommerce/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func ProductCreateValidator(c *fiber.Ctx) error {
	var reqBody dto.CreateProductDto
	if err := c.BodyParser(&reqBody); err != nil {
		return badRequest(c, "Invalid request body")
	}

	err := validate.Struct(reqBody)
	if err != nil {
		// Handle validation errors using the common function
		return handleValidationErrors(c, err, "en")
	}

	return c.Next()
}
