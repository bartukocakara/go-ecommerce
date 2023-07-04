package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler interface {
	GetRoles(c *fiber.Ctx) error
	GetRoleByID(c *fiber.Ctx) error
	CreateRole(c *fiber.Ctx) error
	UpdateRole(c *fiber.Ctx) error
	DeleteRole(c *fiber.Ctx) error
}

type roleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) RoleHandler {
	return &roleHandler{
		roleService: roleService,
	}
}

func (h *roleHandler) GetRoles(c *fiber.Ctx) error {
	Roles, err := h.roleService.GetRoles()
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(Roles)
}

func (h *roleHandler) GetRoleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	RoleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle error
		return err
	}

	Role, err := h.roleService.GetRoleByID(uint(RoleID))
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(Role)
}

func (h *roleHandler) CreateRole(c *fiber.Ctx) error {
	createRoleDTO := new(dto.CreateRoleDTO)
	if err := c.BodyParser(createRoleDTO); err != nil {
		// Handle error
		return err
	}

	Role := &entity.Role{
		FirstName: createRoleDTO.FirstName,
		LastName:  createRoleDTO.LastName,
		Email:     createRoleDTO.Email,
		Password:  createRoleDTO.Password,
	}

	err := h.roleService.CreateRole(Role)
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(Role)
}

func (h *roleHandler) UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")
	RoleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle error
		return err
	}

	updateRoleDTO := new(dto.UpdateRoleDTO)
	if err := c.BodyParser(updateRoleDTO); err != nil {
		// Handle error
		return err
	}

	Role := &entity.Role{
		ID:        uint(RoleID),
		FirstName: updateRoleDTO.FirstName,
		LastName:  updateRoleDTO.LastName,
		Email:     updateRoleDTO.Email,
	}

	err = h.roleService.UpdateRole(Role)
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(Role)
}

func (h *roleHandler) DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")
	RoleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle error
		return err
	}

	Role, err := h.roleService.GetRoleByID(uint(RoleID))
	if err != nil {
		// Handle error
		return err
	}

	err = h.roleService.DeleteRole(Role)
	if err != nil {
		// Handle error
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Role deleted successfully",
	})
}
