package middleware

import "github.com/gofiber/fiber/v2"

func CORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Response().Header.Set("Access-Control-Allow-Origin", "*")
		c.Response().Header.Set("Access-Control-Allow-Headers", "*")
		c.Response().Header.Set("Access-Control-Allow-Methods", "*")
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	}
}
