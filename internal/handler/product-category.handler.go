package handler

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ProductCategoryHandler interface {
	GetProductCategorys(c *fiber.Ctx) error
	GetProductCategoryByID(c *fiber.Ctx) error
	CreateProductCategory(c *fiber.Ctx) error
	UpdateProductCategory(c *fiber.Ctx) error
	DeleteProductCategory(c *fiber.Ctx) error
}

type productCategoryHandler struct {
	productCategoryService service.ProductCategoryService
}

func NewProductCategoryHandler(productCategoryService service.ProductCategoryService) ProductCategoryHandler {
	return &productCategoryHandler{
		productCategoryService: productCategoryService,
	}
}

func (h *productCategoryHandler) GetProductCategorys(c *fiber.Ctx) error {
	page, perPage, err := validateQueryParams(c, "page", "per_page", 1, 10)
	if err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Invalid page number or per_page value")
	}

	// Parse and extract filter parameters from the request query
	filter := &dto.FilterProductCategoryDTO{
		Name: c.Query("name", ""),
	}

	productCategorys, total, err := h.productCategoryService.GetProductCategorys(page, perPage, filter)
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error fetching product-categorys")
	}

	var productCategorysInterfaceSlice []interface{}
	for _, productCategory := range productCategorys {
		productCategorysInterfaceSlice = append(productCategorysInterfaceSlice, productCategory)
	}

	return createPaginatedResponse(c, fiber.StatusOK, "OK", productCategorysInterfaceSlice, page, perPage, total)
}

func (h *productCategoryHandler) GetProductCategoryByID(c *fiber.Ctx) error {
	id := parseIDParam(c)

	productCategory, err := h.productCategoryService.GetProductCategoryByID(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(productCategory)
}

func (h *productCategoryHandler) CreateProductCategory(c *fiber.Ctx) error {
	productCategory := new(entity.ProductCategory)

	if err := parseBody(c, productCategory); err != nil {
		return handleError(c, err)
	}

	err := h.productCategoryService.CreateProductCategory(productCategory)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(productCategory)
}

func (h *productCategoryHandler) UpdateProductCategory(c *fiber.Ctx) error {
	id := parseIDParam(c)
	productCategory := new(entity.ProductCategory)

	if err := parseBody(c, productCategory); err != nil {
		return handleError(c, err)
	}

	productCategory.ID = id

	err := h.productCategoryService.UpdateProductCategory(productCategory)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(productCategory)
}

func (h *productCategoryHandler) DeleteProductCategory(c *fiber.Ctx) error {
	id := parseIDParam(c)

	err := h.productCategoryService.DeleteProductCategory(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
