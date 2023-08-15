package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type ProductCategoryService interface {
	GetProductCategoryByID(id uint) (*entity.ProductCategory, error)
	GetProductCategorys(offset, limit int, filterDto *dto.FilterProductCategoryDTO) ([]*entity.ProductCategory, int, error)
	CreateProductCategory(productCategory *entity.ProductCategory) error
	UpdateProductCategory(productCategory *entity.ProductCategory) error
	DeleteProductCategory(id uint) error
}

type productCategoryService struct {
	productCategoryRepository repository.ProductCategoryRepository
}

func NewProductCategoryService(productCategoryRepository repository.ProductCategoryRepository) ProductCategoryService {
	return &productCategoryService{
		productCategoryRepository: productCategoryRepository,
	}
}

func (s *productCategoryService) GetProductCategorys(offset, limit int, filter *dto.FilterProductCategoryDTO) ([]*entity.ProductCategory, int, error) {
	return s.productCategoryRepository.GetProductCategorys(offset, limit, filter)
}

func (s *productCategoryService) GetProductCategoryByID(id uint) (*entity.ProductCategory, error) {
	productCategory, err := s.productCategoryRepository.GetProductCategoryByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}

	return productCategory, nil
}

func (s *productCategoryService) CreateProductCategory(productCategory *entity.ProductCategory) error {
	err := s.productCategoryRepository.CreateProductCategory(productCategory)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *productCategoryService) UpdateProductCategory(productCategory *entity.ProductCategory) error {
	err := s.productCategoryRepository.UpdateProductCategory(productCategory)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *productCategoryService) DeleteProductCategory(id uint) error {
	err := s.productCategoryRepository.DeleteProductCategory(id)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}
