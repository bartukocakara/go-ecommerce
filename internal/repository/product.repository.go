package repository

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts(offset, limit int, filterDto *dto.FilterProductDto) ([]entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(product *entity.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProductByID(id uint) (*entity.Product, error) {
	var product entity.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepository) GetProducts(offset, limit int, filterDto *dto.FilterProductDto) ([]entity.Product, error) {
	var products []entity.Product

	// Apply pagination and filtering logic
	query := r.db.Offset(offset).Limit(limit)

	if filterDto != nil {
		if filterDto.Name != "" {
			query = query.Where("name LIKE ?", "%"+filterDto.Name+"%")
		}
	}

	result := query.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *productRepository) CreateProduct(product *entity.Product) error {
	result := r.db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *productRepository) UpdateProduct(product *entity.Product) error {
	result := r.db.Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *productRepository) DeleteProduct(product *entity.Product) error {
	result := r.db.Delete(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
