package repository

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	GetProductCategorys(offset, limit int, filter *dto.FilterProductCategoryDTO) ([]*entity.ProductCategory, int, error)
	GetProductCategoryByID(id uint) (*entity.ProductCategory, error)
	CreateProductCategory(productCategory *entity.ProductCategory) error
	UpdateProductCategory(productCategory *entity.ProductCategory) error
	DeleteProductCategory(id uint) error
}

type productCategoryRepository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{
		db: db,
	}
}

func (r *productCategoryRepository) GetProductCategorys(offset, limit int, filter *dto.FilterProductCategoryDTO) ([]*entity.ProductCategory, int, error) {
	var productCategorys []*entity.ProductCategory

	db := r.db.Model(&entity.ProductCategory{})

	if filter != nil {
		if filter.Name != "" {
			db = db.Where("name ILIKE ?", "%"+filter.Name+"%")
		}
	}
	count, err := CountTotal(db, productCategorys)
	if err != nil {
		return nil, 0, err
	}
	db = Paginate(db, offset, limit)

	if err := db.Find(&productCategorys).Error; err != nil {
		return nil, 0, err
	}

	return productCategorys, int(count), nil
}

func (r *productCategoryRepository) GetProductCategoryByID(id uint) (*entity.ProductCategory, error) {
	var productCategory entity.ProductCategory
	result := r.db.First(&productCategory, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &productCategory, nil
}

func (r *productCategoryRepository) CreateProductCategory(productCategory *entity.ProductCategory) error {
	result := r.db.Create(productCategory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *productCategoryRepository) UpdateProductCategory(productCategory *entity.ProductCategory) error {
	result := r.db.Save(productCategory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *productCategoryRepository) DeleteProductCategory(id uint) error {
	result := r.db.Delete(id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
