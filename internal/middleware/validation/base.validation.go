package validator

import (
	"ecommerce/internal/validation"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func handleValidationErrors(ctx *fiber.Ctx, err error, lang string) error {
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return badRequest(ctx, "Validation failed")
	}

	var validationErrors []validationError
	for _, e := range errors {
		message := fmt.Sprintf("Validation failed on '%s' tag", e.Tag())

		// Get the validation message based on the language and tag
		if langMessages, ok := validation.Messages[lang]; ok {
			if msg, ok := langMessages[e.Tag()]; ok {
				message = msg
			}
		}

		// Replace :field, :min, and :max with actual values
		message = strings.ReplaceAll(message, ":field", e.Field())
		message = strings.ReplaceAll(message, ":min", e.Param())
		message = strings.ReplaceAll(message, ":max", e.Param())

		validationErrors = append(validationErrors, validationError{
			Field:   e.Field(),
			Message: message,
		})
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"statusCode": fiber.StatusUnprocessableEntity,
		"message":    "Unprocessable Content",
		"errors":     validationErrors,
	})
}

func badRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"statusCode": fiber.StatusBadRequest,
		"message":    message,
	})
}
