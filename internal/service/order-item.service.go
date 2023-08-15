package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type OrderItemService interface {
	GetOrderItemByID(id uint) (*entity.OrderItem, error)
	GetOrderItems(offset, limit int, filterDto *dto.FilterOrderItemDTO) ([]*entity.OrderItem, int, error)
	CreateOrderItem(orderItem *entity.OrderItem) error
	UpdateOrderItem(orderItem *entity.OrderItem) error
	DeleteOrderItem(id uint) error
}

type orderItemService struct {
	orderItemRepository repository.OrderItemRepository
}

func NewOrderItemService(orderItemRepository repository.OrderItemRepository) OrderItemService {
	return &orderItemService{
		orderItemRepository: orderItemRepository,
	}
}

func (s *orderItemService) GetOrderItems(offset, limit int, filter *dto.FilterOrderItemDTO) ([]*entity.OrderItem, int, error) {
	return s.orderItemRepository.GetOrderItems(offset, limit, filter)
}

func (s *orderItemService) GetOrderItemByID(id uint) (*entity.OrderItem, error) {
	orderItem, err := s.orderItemRepository.GetOrderItemByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}

	return orderItem, nil
}

func (s *orderItemService) CreateOrderItem(orderItem *entity.OrderItem) error {
	err := s.orderItemRepository.CreateOrderItem(orderItem)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *orderItemService) UpdateOrderItem(orderItem *entity.OrderItem) error {
	err := s.orderItemRepository.UpdateOrderItem(orderItem)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *orderItemService) DeleteOrderItem(id uint) error {
	err := s.orderItemRepository.DeleteOrderItem(id)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}
