package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/enums"
	"ecommerce/internal/service"
	"fmt"

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
		fmt.Println("Error parsing request body:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
	// Call the corresponding service method to handle the registration logic
	registerResponse, err := h.AuthService.Register(registerDto)
	if err != nil {
		return createErrorResponse(ctx, fiber.StatusBadRequest, "Failed to register")
	}

	// Return the custom response
	response := CreateTokenResponse(fiber.StatusCreated, "OK", map[string]interface{}{
		"user":         registerResponse.User,
		"role":         registerResponse.Role,
		"access_token": registerResponse.AccessToken,
	})

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var loginDto dto.LoginDto
	if err := ctx.BodyParser(&loginDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	// Call the corresponding service method to handle the login logic
	loginResponse, err := h.AuthService.Login(loginDto)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Return the custom response
	response := CreateTokenResponse(fiber.StatusOK, "OK", map[string]interface{}{
		"user":         loginResponse.User,
		"role":         loginResponse.Role,
		"access_token": loginResponse.AccessToken,
	})

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthHandler) ForgotPassword(ctx *fiber.Ctx) error {
	var forgotPasswordDto dto.ForgotPasswordDto
	if err := ctx.BodyParser(&forgotPasswordDto); err != nil {
		return err
	}

	response, result, err := h.AuthService.ForgotPassword(forgotPasswordDto)

	if err != nil {
		return createErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to process forget password request")
	}

	// Handle different resulting cases
	switch result {
	case enums.UserNotFound:
		return createErrorResponse(ctx, fiber.StatusNotFound, response)

	case enums.TokenAlreadyExists:
		response := CreateResponse(fiber.StatusBadRequest, response, fiber.Map{})
		return ctx.Status(fiber.StatusBadRequest).JSON(response)

	case enums.Success:
		response := CreateResponse(fiber.StatusCreated, response, fiber.Map{})
		return ctx.Status(fiber.StatusCreated).JSON(response)
	default:
		return createErrorResponse(ctx, fiber.StatusBadRequest, "Unknown error occurred")
	}

}
