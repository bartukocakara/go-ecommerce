package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the JWT token from the Authorization header
		authHeader := c.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide the secret key or public key to validate the token
			// Example: return []byte("secret-key"), nil
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			// Return an error response if the token is invalid
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired JWT token")
		}

		// Set the authenticated user ID in the context
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64)
		email := claims["email"].(string)
		c.Locals("user_id", userID)
		userRole := claims["role"].(string) // Replace "role" with the actual key in the claims
		c.Locals("userRole", userRole)
		c.Locals("email", email)

		// Continue to the next middleware or route handler
		return c.Next()
	}
}
