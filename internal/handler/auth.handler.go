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
