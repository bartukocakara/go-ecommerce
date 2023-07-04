package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(users)
}

func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle error
		return err
	}

	user, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(user)
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	createUserDTO := new(dto.CreateUserDTO)
	if err := c.BodyParser(createUserDTO); err != nil {
		// Handle error
		return err
	}

	user := &entity.User{
		FirstName: createUserDTO.FirstName,
		LastName:  createUserDTO.LastName,
		Email:     createUserDTO.Email,
		Password:  createUserDTO.Password,
	}

	err := h.userService.CreateUser(user)
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(user)
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle error
		return err
	}

	updateUserDTO := new(dto.UpdateUserDTO)
	if err := c.BodyParser(updateUserDTO); err != nil {
		// Handle error
		return err
	}

	user := &entity.User{
		ID:        uint(userID),
		FirstName: updateUserDTO.FirstName,
		LastName:  updateUserDTO.LastName,
		Email:     updateUserDTO.Email,
	}

	err = h.userService.UpdateUser(user)
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(user)
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle error
		return err
	}

	user, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		// Handle error
		return err
	}

	err = h.userService.DeleteUser(user)
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}