package repository

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type BasketRepository interface {
	GetBaskets() ([]entity.Basket, error)
	GetBasketByID(id uint) (*entity.Basket, error)
	CreateBasket(basket *entity.Basket) error
	UpdateBasket(basket *entity.Basket) error
	DeleteBasket(basket *entity.Basket) error
}

type basketRepository struct {
	db *gorm.DB
}

func NewBasketRepository(db *gorm.DB) BasketRepository {
	return &basketRepository{
		db: db,
	}
}

func (r *basketRepository) GetBaskets() ([]entity.Basket, error) {
	var baskets []entity.Basket
	result := r.db.Find(&baskets)
	if result.Error != nil {
		return nil, result.Error
	}
	return baskets, nil
}

func (r *basketRepository) GetBasketByID(id uint) (*entity.Basket, error) {
	var basket entity.Basket
	result := r.db.First(&basket, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &basket, nil
}

func (r *basketRepository) CreateBasket(basket *entity.Basket) error {
	result := r.db.Create(basket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *basketRepository) UpdateBasket(basket *entity.Basket) error {
	result := r.db.Save(basket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *basketRepository) DeleteBasket(basket *entity.Basket) error {
	result := r.db.Delete(basket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Implement the methods of the BasketRepository interface...
