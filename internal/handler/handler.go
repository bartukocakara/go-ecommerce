package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func parsePaginationParams(c *fiber.Ctx) (int, int) {
	offsetParam := c.Query("offset") // Assuming "offset" is the query parameter for offset
	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		offset = 0 // Default value if the query parameter is not provided or invalid
	}

	limitParam := c.Query("limit") // Assuming "limit" is the query parameter for limit
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10 // Default value if the query parameter is not provided or invalid
	}

	return offset, limit
}

func parseFilterParams(c *fiber.Ctx) interface{} {
	// filterValue := c.Query("filter") // Assuming "filter" is the query parameter for filter

	// Implement your logic to parse the filter parameter based on your requirements
	// and return the appropriate filter DTO or data structure
	// Example:
	/*
		filterDto := &dto.FilterDto{
			FilterValue: filterValue,
		}

		return filterDto
	*/

	return nil // Return nil if no filter is applied
}

func parseIDParam(c *fiber.Ctx) uint {
	idParam := c.Params("id") // Assuming "id" is the route parameter for ID

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		// Handle the error or return an appropriate response
	}

	return uint(id)
}

func parseBody(c *fiber.Ctx, entity interface{}) error {
	// Implement your logic to parse the request body and populate the entity struct
	// Return any error that occurred during the parsing process
	if err := c.BodyParser(entity); err != nil {
		// Handle the error or return an appropriate response
		return err
	}

	return nil
}

func handleError(c *fiber.Ctx, err error) error {
	// Implement your logic to handle and format the error response
	// Return the appropriate error response based on the provided error
	return nil
}
