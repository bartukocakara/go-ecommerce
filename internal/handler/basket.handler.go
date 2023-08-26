package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BasketHandler interface {
	List(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type basketHandler struct {
	basketService service.BasketService
}

func NewBasketHandler(basketService service.BasketService) BasketHandler {
	return &basketHandler{
		basketService: basketService,
	}
}

func (h *basketHandler) List(c *fiber.Ctx) error {
	baskets, err := h.basketService.List()
	if err != nil {
		return err
	}

	return c.JSON(baskets)
}

func (h *basketHandler) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	Basket, err := h.basketService.Show(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketHandler) Create(c *fiber.Ctx) error {
	var createDto dto.CreateBasketDto
	err := c.BodyParser(&createDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketService.Create(&createDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	var updateDto dto.UpdateBasketDto
	err = c.BodyParser(&updateDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketService.Update(uint(id), &updateDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	err = h.basketService.Delete(uint(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
