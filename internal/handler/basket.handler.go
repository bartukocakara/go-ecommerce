package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BasketHandler interface {
	GetBaskets(c *fiber.Ctx) error
	GetBasketByID(c *fiber.Ctx) error
	CreateBasket(c *fiber.Ctx) error
	UpdateBasket(c *fiber.Ctx) error
	DeleteBasket(c *fiber.Ctx) error
}

type basketHandler struct {
	basketService service.BasketService
}

func NewBasketHandler(basketService service.BasketService) BasketHandler {
	return &basketHandler{
		basketService: basketService,
	}
}

func (h *basketHandler) GetBaskets(c *fiber.Ctx) error {
	baskets, err := h.basketService.GetBaskets()
	if err != nil {
		return err
	}

	return c.JSON(baskets)
}

func (h *basketHandler) GetBasketByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	Basket, err := h.basketService.GetBasketByID(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketHandler) CreateBasket(c *fiber.Ctx) error {
	var createDto dto.CreateBasketDto
	err := c.BodyParser(&createDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketService.CreateBasket(&createDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketHandler) UpdateBasket(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	var updateDto dto.UpdateBasketDto
	err = c.BodyParser(&updateDto)
	if err != nil {
		return err
	}

	Basket, err := h.basketService.UpdateBasket(uint(id), &updateDto)
	if err != nil {
		return err
	}

	return c.JSON(Basket)
}

func (h *basketHandler) DeleteBasket(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return err
	}

	err = h.basketService.DeleteBasket(uint(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
