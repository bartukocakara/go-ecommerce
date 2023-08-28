package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	List(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

const (
	RoleAdmin    = 1
	RoleCustomer = 2
)

func (h *userHandler) List(c *fiber.Ctx) error {
	page, perPage, err := validateQueryParams(c, "page", "per_page", 1, 10)
	if err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Invalid page number or per_page value")
	}

	// Parse and extract filter parameters from the request query
	filter := &dto.FilterUserDTO{
		FirstName: c.Query("first_name", ""),
		LastName:  c.Query("last_name", ""),
		Email:     c.Query("email", ""),
	}
	users, total, err := h.userService.List(page, perPage, filter)
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error fetching users")
	}

	var usersInterfaceSlice []interface{}
	for _, user := range users {
		usersInterfaceSlice = append(usersInterfaceSlice, user)
	}

	return createPaginatedResponse(c, fiber.StatusOK, "OK", usersInterfaceSlice, page, perPage, total)
}

func (h *userHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle error
		return err
	}

	user, err := h.userService.Show(uint(userID))
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error fetching user")
	}
	response := CreateResponse(fiber.StatusOK, "OK", user)
	return c.Status(fiber.StatusBadRequest).JSON(response)
}

func (h *userHandler) Create(c *fiber.Ctx) error {
	createUserDTO := new(dto.CreateUserDTO)
	if err := c.BodyParser(createUserDTO); err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error creating user")
	}

	user, err := h.userService.Create(createUserDTO)
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error creating user")
	}

	response := CreateResponse(fiber.StatusCreated, "Created", user)
	return c.Status(fiber.StatusBadRequest).JSON(response)
}

func (h *userHandler) Update(c *fiber.Ctx) error {
	id := parseIDParam(c)

	updateUserDTO := new(dto.UpdateUserDTO)
	if err := c.BodyParser(updateUserDTO); err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Error parsing body")
	}

	err := h.userService.Update(id, updateUserDTO)
	if err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Error updating user")
	}

	response := CreateResponse(fiber.StatusNoContent, "Updated", err)
	return c.Status(fiber.StatusNoContent).JSON(response)
}

func (h *userHandler) Delete(c *fiber.Ctx) error {
	id := parseIDParam(c)

	err := h.userService.Delete(id)
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	response := CreateResponse(fiber.StatusNoContent, "Delete", err)
	return c.Status(fiber.StatusNoContent).JSON(response)

}
