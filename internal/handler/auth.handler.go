package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	var registerDto dto.RegisterDto
	if err := ctx.BodyParser(&registerDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	// Call the corresponding service method to handle the registration logic
	err := h.AuthService.Register(registerDto)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to register")
	}

	return ctx.JSON(fiber.Map{"message": "Registration successful"})
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var loginDto dto.LoginDto
	if err := ctx.BodyParser(&loginDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	// Call the corresponding service method to handle the login logic
	token, err := h.AuthService.Login(loginDto)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return ctx.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) ForgetPassword(c *fiber.Ctx) error {
	// Parse the request body into a DTO
	var forgetPasswordDto dto.ForgetPasswordDto
	if err := c.BodyParser(&forgetPasswordDto); err != nil {
		return err
	}

	// Call the service method
	resetToken, err := h.AuthService.ForgetPassword(forgetPasswordDto)

	if err != nil {
		// Handle the error and return an appropriate response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to process forget password request",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token":   resetToken,
		"message": "Forget password request processed successfully",
	})
}
