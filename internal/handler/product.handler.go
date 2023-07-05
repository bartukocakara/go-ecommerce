package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("perPage", "10"))
	filterDto := &dto.FilterProductDto{
		Name: c.Query("name"),
	}

	products, err := h.productService.GetProducts(page, perPage, filterDto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve products",
			"error":   err.Error(),
		})
	}

	return c.JSON(products)
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	createDto := new(dto.CreateProductDto)
	if err := c.BodyParser(createDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	err := h.productService.CreateProduct(createDto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product created successfully",
	})
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	updateDto := new(dto.UpdateProductDto)
	if err := c.BodyParser(updateDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	err := h.productService.UpdateProduct(uint(id), updateDto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := h.productService.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
