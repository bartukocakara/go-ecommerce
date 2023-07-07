package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type BasketItemService interface {
	GetBasketItems() ([]entity.BasketItem, error)
	GetBasketItemByID(id uint) (*entity.BasketItem, error)
	CreateBasketItem(createDto *dto.CreateBasketItemDto) (*entity.BasketItem, error)
	UpdateBasketItem(id uint, basketItemDto *dto.UpdateBasketItemDto) (*entity.BasketItem, error)
	DeleteBasketItem(id uint) error
}

type basketItemService struct {
	basketItemRepository repository.BasketItemRepository
}

func NewBasketItemService(basketItemRepository repository.BasketItemRepository) BasketItemService {
	return &basketItemService{
		basketItemRepository: basketItemRepository,
	}
}

func (s *basketItemService) GetBasketItems() ([]entity.BasketItem, error) {
	return s.basketItemRepository.GetBasketItems()
}

func (s *basketItemService) GetBasketItemByID(id uint) (*entity.BasketItem, error) {
	return s.basketItemRepository.GetBasketItemByID(id)
}

func (s *basketItemService) CreateBasketItem(createDto *dto.CreateBasketItemDto) (*entity.BasketItem, error) {
	basketItem := &entity.BasketItem{
		Quantity: createDto.Quantity,
	}

	err := s.basketItemRepository.CreateBasketItem(basketItem)
	if err != nil {
		return nil, err
	}

	return basketItem, nil
}

func (s *basketItemService) UpdateBasketItem(id uint, updateDto *dto.UpdateBasketItemDto) (*entity.BasketItem, error) {
	basketItem, err := s.basketItemRepository.GetBasketItemByID(id)
	if err != nil {
		return nil, err
	}

	basketItem.Quantity = updateDto.Quantity

	err = s.basketItemRepository.UpdateBasketItem(basketItem)
	if err != nil {
		return nil, err
	}

	return basketItem, nil
}

func (s *basketItemService) DeleteBasketItem(id uint) error {
	basketItem, err := s.basketItemRepository.GetBasketItemByID(id)
	if err != nil {
		return err
	}

	return s.basketItemRepository.DeleteBasketItem(basketItem)
}

// Implement the methods of the BasketItemService interface...
