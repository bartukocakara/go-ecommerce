package repository

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type BasketItemRepository interface {
	List() ([]entity.BasketItem, error)
	Show(id uint) (*entity.BasketItem, error)
	Create(basketItem *entity.BasketItem) error
	Update(basketItem *entity.BasketItem) error
	Delete(basketItem *entity.BasketItem) error
}

type basketItemRepository struct {
	db *gorm.DB
}

func NewBasketItemRepository(db *gorm.DB) BasketItemRepository {
	return &basketItemRepository{
		db: db,
	}
}

func (r *basketItemRepository) List() ([]entity.BasketItem, error) {
	var basketItems []entity.BasketItem
	result := r.db.Find(&basketItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return basketItems, nil
}

func (r *basketItemRepository) Show(id uint) (*entity.BasketItem, error) {
	var basketItem entity.BasketItem
	result := r.db.First(&basketItem, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &basketItem, nil
}

func (r *basketItemRepository) Create(basketItem *entity.BasketItem) error {
	result := r.db.Create(basketItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *basketItemRepository) Update(basketItem *entity.BasketItem) error {
	result := r.db.Save(basketItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *basketItemRepository) Delete(basketItem *entity.BasketItem) error {
	result := r.db.Delete(basketItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Implement the methods of the BasketRepository interface...
