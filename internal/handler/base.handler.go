package handler

import (
	"ecommerce/internal/utils"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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

func createResponse(statusCode int, message string, data fiber.Map) fiber.Map {
	return fiber.Map{
		"message":    message,
		"statusCode": statusCode,
		"status":     getStatusText(statusCode),
		"result":     data,
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

	response := createResponse(statusCode, message, fiber.Map{
		"data":       data,
		"pagination": pagination,
	})

	return c.Status(statusCode).JSON(response)
}

func createErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(utils.GenericResponse{
		Message:   message,
		StatusCode: status,
		Status:     "error",
		Result: utils.GenericResult{
			Data:       nil,
			Pagination: utils.Pagination{},
		},
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