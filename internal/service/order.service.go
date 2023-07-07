package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type OrderService interface {
	GetOrderByID(id uint) (*entity.Order, error)
	GetOrders(offset, limit int, filterDto *dto.FilterOrderDto) ([]*entity.Order, error)
	CreateOrder(order *entity.Order) error
	UpdateOrder(order *entity.Order) error
	DeleteOrder(id uint) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (s *orderService) GetOrderByID(id uint) (*entity.Order, error) {
	return s.orderRepository.GetOrderByID(id)
}

func (s *orderService) GetOrders(offset, limit int, filterOrderDto *dto.FilterOrderDto) ([]*entity.Order, error) {
	return s.orderRepository.GetOrders(offset, limit, filterOrderDto)
}

func (s *orderService) CreateOrder(order *entity.Order) error {
	return s.orderRepository.CreateOrder(order)
}

func (s *orderService) UpdateOrder(order *entity.Order) error {
	return s.orderRepository.UpdateOrder(order)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.orderRepository.DeleteOrder(id)
}
