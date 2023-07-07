package repository

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrders(offset, limit int, filterDto *dto.FilterOrderDto) ([]*entity.Order, error)
	GetOrderByID(id uint) (*entity.Order, error)
	CreateOrder(order *entity.Order) error
	UpdateOrder(order *entity.Order) error
	DeleteOrder(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) GetOrders(offset, limit int, filterDto *dto.FilterOrderDto) ([]*entity.Order, error) {
	var orders []*entity.Order

	query := r.db.Model(&entity.Order{})

	if filterDto != nil {
		if filterDto.Name != "" {
			query = query.Where("name LIKE ?", "%"+filterDto.Name+"%")
		}
	}

	query = paginateQuery(query, offset, limit)

	result := query.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (r *orderRepository) GetOrderByID(id uint) (*entity.Order, error) {
	var order entity.Order
	result := r.db.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *orderRepository) CreateOrder(order *entity.Order) error {
	result := r.db.Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepository) UpdateOrder(order *entity.Order) error {
	result := r.db.Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepository) DeleteOrder(id uint) error {
	result := r.db.Delete(&entity.Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
