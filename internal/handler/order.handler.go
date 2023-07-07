package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/service"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx) error
	GetOrderByID(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
	DeleteOrder(c *fiber.Ctx) error
}

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) OrderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (h *orderHandler) GetOrders(c *fiber.Ctx) error {
	offset, limit := parsePaginationParams(c)
	filterOrderDto := parseFilterParams(c).(*dto.FilterOrderDto) // Perform type assertion

	orders, err := h.orderService.GetOrders(offset, limit, filterOrderDto)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(orders)
}

func (h *orderHandler) GetOrderByID(c *fiber.Ctx) error {
	id := parseIDParam(c)

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(order)
}

func (h *orderHandler) CreateOrder(c *fiber.Ctx) error {
	order := new(entity.Order)

	if err := parseBody(c, order); err != nil {
		return handleError(c, err)
	}

	err := h.orderService.CreateOrder(order)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(order)
}

func (h *orderHandler) UpdateOrder(c *fiber.Ctx) error {
	id := parseIDParam(c)
	order := new(entity.Order)

	if err := parseBody(c, order); err != nil {
		return handleError(c, err)
	}

	order.ID = id

	err := h.orderService.UpdateOrder(order)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(order)
}

func (h *orderHandler) DeleteOrder(c *fiber.Ctx) error {
	id := parseIDParam(c)

	err := h.orderService.DeleteOrder(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
