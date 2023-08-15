package repository

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	GetOrderItems(offset, limit int, filter *dto.FilterOrderItemDTO) ([]*entity.OrderItem, int, error)
	GetOrderItemByID(id uint) (*entity.OrderItem, error)
	CreateOrderItem(orderItem *entity.OrderItem) error
	UpdateOrderItem(orderItem *entity.OrderItem) error
	DeleteOrderItem(id uint) error
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{
		db: db,
	}
}

func (r *orderItemRepository) GetOrderItems(offset, limit int, filter *dto.FilterOrderItemDTO) ([]*entity.OrderItem, int, error) {
	var orderItems []*entity.OrderItem

	db := r.db.Model(&entity.OrderItem{})

	if filter != nil {
		if filter.Name != "" {
			db = db.Where("name ILIKE ?", "%"+filter.Name+"%")
		}
	}
	count, err := CountTotal(db, orderItems)
	if err != nil {
		return nil, 0, err
	}
	db = Paginate(db, offset, limit)

	if err := db.Find(&orderItems).Error; err != nil {
		return nil, 0, err
	}

	return orderItems, int(count), nil
}

func (r *orderItemRepository) GetOrderItemByID(id uint) (*entity.OrderItem, error) {
	var orderItem entity.OrderItem
	result := r.db.First(&orderItem, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orderItem, nil
}

func (r *orderItemRepository) CreateOrderItem(orderItem *entity.OrderItem) error {
	result := r.db.Create(orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderItemRepository) UpdateOrderItem(orderItem *entity.OrderItem) error {
	result := r.db.Save(orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderItemRepository) DeleteOrderItem(id uint) error {
	result := r.db.Delete(id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
