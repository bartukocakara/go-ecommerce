package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type ProductService interface {
	GetProducts(page, perPage int, filterDto *dto.FilterProductDto) ([]entity.Product, error)
	CreateProduct(createDto *dto.CreateProductDto) error
	UpdateProduct(id uint, updateDto *dto.UpdateProductDto) error
	DeleteProduct(id uint) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) GetProducts(page, perPage int, filterDto *dto.FilterProductDto) ([]entity.Product, error) {
	// Calculate offset and limit for pagination
	offset := (page - 1) * perPage
	limit := perPage

	products, err := s.productRepository.GetProducts(offset, limit, filterDto)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) CreateProduct(createDto *dto.CreateProductDto) error {
	product := &entity.Product{
		Name:  createDto.Name,
		Price: createDto.Price,
	}

	err := s.productRepository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *productService) UpdateProduct(id uint, updateDto *dto.UpdateProductDto) error {
	product, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return err
	}

	product.Name = updateDto.Name
	product.Price = updateDto.Price

	err = s.productRepository.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *productService) DeleteProduct(id uint) error {
	product, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return err
	}

	err = s.productRepository.DeleteProduct(product)
	if err != nil {
		return err
	}

	return nil
}
