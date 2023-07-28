package response

import "github.com/gofiber/fiber/v2"

func CreateResponse(status int, message string, data fiber.Map) fiber.Map {
	return fiber.Map{
		"statusCode": status,
		"message":    message,
		"result":     data,
	}
}

type GenericErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
