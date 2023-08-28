package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BasketItemHandler interface {
	List(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type basketItemHandler struct {
	basketItemService service.BasketItemService
}

func NewBasketItemHandler(basketItemService service.BasketItemService) BasketItemHandler {
	return &basketItemHandler{
		basketItemService: basketItemService,
	}
}

func (h *basketItemHandler) List(c *fiber.Ctx) error {
	basketItems, err := h.basketItemService.List()
	if err != nil {
		return err
	}

	return c.JSON(basketItems)
}

func (h *basketItemHandler) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	Basket, err := h.basketItemService.Show(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketItemHandler) Create(c *fiber.Ctx) error {
	var createDto dto.CreateBasketItemDto
	err := c.BodyParser(&createDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketItemService.Create(&createDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketItemHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	var updateDto dto.UpdateBasketItemDto
	err = c.BodyParser(&updateDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketItemService.Update(uint(id), &updateDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketItemHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	err = h.basketItemService.Delete(uint(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
