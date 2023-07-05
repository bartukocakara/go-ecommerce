package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler interface {
	GetCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

func (h *categoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		return err
	}

	return c.JSON(categories)
}

func (h *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	category, err := h.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(category)
}

func (h *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	var createDto dto.CreateCategoryDto
	err := c.BodyParser(&createDto)
	if err != nil {
		return err
	}

	category, err := h.categoryService.CreateCategory(createDto)
	if err != nil {
		return err
	}

	return c.JSON(category)
}

func (h *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	var updateDto dto.UpdateCategoryDto
	err = c.BodyParser(&updateDto)
	if err != nil {
		return err
	}

	category, err := h.categoryService.UpdateCategory(uint(id), updateDto)
	if err != nil {
		return err
	}

	return c.JSON(category)
}

func (h *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	err = h.categoryService.DeleteCategory(uint(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
