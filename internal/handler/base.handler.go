package handler

import (
	"ecommerce/internal/handler/response"
	"ecommerce/internal/utils"
	"ecommerce/internal/validation"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type OkResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Result     interface{} `json:"result"`
}

func createOkResponse(statusCode int, message string) fiber.Map {
	return fiber.Map{
		"message":    message,
		"statusCode": statusCode,
		"status":     getStatusFromCode(statusCode),
		"result":     nil,
	}
}

func getStatusFromCode(statusCode int) string {
	if statusCode >= 200 && statusCode <= 299 {
		return "success"
	}
	return "error"
}

func CreateResponse(statusCode int, message string, data fiber.Map) fiber.Map {
	return fiber.Map{
		"message":    message,
		"statusCode": statusCode,
		"status":     getStatusText(statusCode),
		"result":     data,
	}
}

func CreateTokenResponse(statusCode int, message string, data map[string]interface{}) fiber.Map {
	return fiber.Map{
		"message":    message,
		"statusCode": statusCode,
		"status":     getStatusText(statusCode),
		"result": fiber.Map{
			"user":         data["user"],
			"role":         data["role"],
			"access_token": data["access_token"],
			"token_type":   "Bearer",
			"expires_in":   3600,
		},
	}
}

func getStatusText(code int) string {
	switch code {
	case fiber.StatusOK:
		return "success"
	case fiber.StatusCreated:
		return "created"
	case fiber.StatusBadRequest:
		return "bad request"
	case fiber.StatusUnauthorized:
		return "unauthorized"
	case fiber.StatusInternalServerError:
		return "internal server error"
	default:
		return "unknown"
	}
}

func createPaginatedResponse(c *fiber.Ctx, statusCode int, message string, data interface{}, currentPage, perPage, total int) error {
	sliceData, ok := data.([]interface{})
	if !ok {
		// Return an error or handle the incorrect type gracefully
		return fmt.Errorf("invalid data type, expected []interface{}")
	}

	pagination := utils.Pagination{
		CurrentPage: currentPage,
		From:        (currentPage - 1) * perPage,
		LastPage:    int(math.Ceil(float64(total) / float64(perPage))),
		PerPage:     perPage,
		To:          (currentPage-1)*perPage + len(sliceData),
		Total:       total,
	}

	response := CreateResponse(statusCode, message, fiber.Map{
		"data":       data,
		"pagination": pagination,
	})

	return c.Status(statusCode).JSON(response)
}

func createErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.JSON(response.GenericErrorResponse{
		Message:    message,
		StatusCode: status,
	})
}

func validateQueryParams(c *fiber.Ctx, pageKey, perPageKey string, defaultPage, defaultPerPage int) (int, int, error) {
	pageStr := c.Query(pageKey, strconv.Itoa(defaultPage))
	perPageStr := c.Query(perPageKey, strconv.Itoa(defaultPerPage))

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, err
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		return 0, 0, err
	}

	return page, perPage, nil
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
