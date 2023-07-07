package repository

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type BasketItemRepository interface {
	GetBasketItems() ([]entity.BasketItem, error)
	GetBasketItemByID(id uint) (*entity.BasketItem, error)
	CreateBasketItem(basketItem *entity.BasketItem) error
	UpdateBasketItem(basketItem *entity.BasketItem) error
	DeleteBasketItem(basketItem *entity.BasketItem) error
}

type basketItemRepository struct {
	db *gorm.DB
}

func NewBasketItemRepository(db *gorm.DB) BasketItemRepository {
	return &basketItemRepository{
		db: db,
	}
}

func (r *basketItemRepository) GetBasketItems() ([]entity.BasketItem, error) {
	var basketItems []entity.BasketItem
	result := r.db.Find(&basketItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return basketItems, nil
}

func (r *basketItemRepository) GetBasketItemByID(id uint) (*entity.BasketItem, error) {
	var basketItem entity.BasketItem
	result := r.db.First(&basketItem, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &basketItem, nil
}

func (r *basketItemRepository) CreateBasketItem(basketItem *entity.BasketItem) error {
	result := r.db.Create(basketItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *basketItemRepository) UpdateBasketItem(basketItem *entity.BasketItem) error {
	result := r.db.Save(basketItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *basketItemRepository) DeleteBasketItem(basketItem *entity.BasketItem) error {
	result := r.db.Delete(basketItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Implement the methods of the BasketRepository interface...
