package repository

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategories() ([]entity.Category, error)
	GetCategoryByID(id uint) (*entity.Category, error)
	CreateCategory(category *entity.Category) error
	UpdateCategory(category *entity.Category) error
	DeleteCategory(category *entity.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetCategories() ([]entity.Category, error) {
	var categories []entity.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *categoryRepository) GetCategoryByID(id uint) (*entity.Category, error) {
	var category entity.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) CreateCategory(category *entity.Category) error {
	result := r.db.Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *categoryRepository) UpdateCategory(category *entity.Category) error {
	result := r.db.Save(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *categoryRepository) DeleteCategory(category *entity.Category) error {
	result := r.db.Delete(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
