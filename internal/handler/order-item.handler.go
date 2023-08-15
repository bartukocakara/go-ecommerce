package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OrderItemHandler interface {
	GetOrderItems(c *fiber.Ctx) error
	GetOrderItemByID(c *fiber.Ctx) error
	CreateOrderItem(c *fiber.Ctx) error
	UpdateOrderItem(c *fiber.Ctx) error
	DeleteOrderItem(c *fiber.Ctx) error
}

type orderItemHandler struct {
	orderItemService service.OrderItemService
}

func NewOrderItemHandler(orderItemService service.OrderItemService) OrderItemHandler {
	return &orderItemHandler{
		orderItemService: orderItemService,
	}
}

func (h *orderItemHandler) GetOrderItems(c *fiber.Ctx) error {
	page, perPage, err := validateQueryParams(c, "page", "per_page", 1, 10)
	if err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Invalid page number or per_page value")
	}

	// Parse and extract filter parameters from the request query
	filter := &dto.FilterOrderItemDTO{
		Name: c.Query("name", ""),
	}

	orderItems, total, err := h.orderItemService.GetOrderItems(page, perPage, filter)
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error fetching order-items")
	}

	var orderItemsInterfaceSlice []interface{}
	for _, orderItem := range orderItems {
		orderItemsInterfaceSlice = append(orderItemsInterfaceSlice, orderItem)
	}

	return createPaginatedResponse(c, fiber.StatusOK, "OK", orderItemsInterfaceSlice, page, perPage, total)
}

func (h *orderItemHandler) GetOrderItemByID(c *fiber.Ctx) error {
	id := parseIDParam(c)

	orderItem, err := h.orderItemService.GetOrderItemByID(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(orderItem)
}

func (h *orderItemHandler) CreateOrderItem(c *fiber.Ctx) error {
	orderItem := new(entity.OrderItem)

	if err := parseBody(c, orderItem); err != nil {
		return handleError(c, err)
	}

	err := h.orderItemService.CreateOrderItem(orderItem)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(orderItem)
}

func (h *orderItemHandler) UpdateOrderItem(c *fiber.Ctx) error {
	id := parseIDParam(c)
	orderItem := new(entity.OrderItem)

	if err := parseBody(c, orderItem); err != nil {
		return handleError(c, err)
	}

	orderItem.ID = id

	err := h.orderItemService.UpdateOrderItem(orderItem)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(orderItem)
}

func (h *orderItemHandler) DeleteOrderItem(c *fiber.Ctx) error {
	id := parseIDParam(c)

	err := h.orderItemService.DeleteOrderItem(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
