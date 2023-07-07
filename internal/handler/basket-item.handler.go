package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BasketItemHandler interface {
	GetBasketItems(c *fiber.Ctx) error
	GetBasketItemByID(c *fiber.Ctx) error
	CreateBasketItem(c *fiber.Ctx) error
	UpdateBasketItem(c *fiber.Ctx) error
	DeleteBasketItem(c *fiber.Ctx) error
}

type basketItemHandler struct {
	basketItemService service.BasketItemService
}

func NewBasketItemHandler(basketItemService service.BasketItemService) BasketItemHandler {
	return &basketItemHandler{
		basketItemService: basketItemService,
	}
}

func (h *basketItemHandler) GetBasketItems(c *fiber.Ctx) error {
	basketItems, err := h.basketItemService.GetBasketItems()
	if err != nil {
		return err
	}

	return c.JSON(basketItems)
}

func (h *basketItemHandler) GetBasketItemByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	Basket, err := h.basketItemService.GetBasketItemByID(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketItemHandler) CreateBasketItem(c *fiber.Ctx) error {
	var createDto dto.CreateBasketItemDto
	err := c.BodyParser(&createDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketItemService.CreateBasketItem(&createDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketItemHandler) UpdateBasketItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	var updateDto dto.UpdateBasketItemDto
	err = c.BodyParser(&updateDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketItemService.UpdateBasketItem(uint(id), &updateDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketItemHandler) DeleteBasketItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	err = h.basketItemService.DeleteBasketItem(uint(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
