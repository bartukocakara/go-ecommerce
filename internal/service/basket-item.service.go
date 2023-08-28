package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type BasketItemService interface {
	List() ([]entity.BasketItem, error)
	Show(id uint) (*entity.BasketItem, error)
	Create(createDto *dto.CreateBasketItemDto) (*entity.BasketItem, error)
	Update(id uint, basketItemDto *dto.UpdateBasketItemDto) (*entity.BasketItem, error)
	Delete(id uint) error
}

type basketItemService struct {
	basketItemRepository repository.BasketItemRepository
}

func NewBasketItemService(basketItemRepository repository.BasketItemRepository) BasketItemService {
	return &basketItemService{
		basketItemRepository: basketItemRepository,
	}
}

func (s *basketItemService) List() ([]entity.BasketItem, error) {
	return s.basketItemRepository.List()
}

func (s *basketItemService) Show(id uint) (*entity.BasketItem, error) {
	return s.basketItemRepository.Show(id)
}

func (s *basketItemService) Create(createDto *dto.CreateBasketItemDto) (*entity.BasketItem, error) {
	basketItem := &entity.BasketItem{
		Quantity: createDto.Quantity,
	}

	err := s.basketItemRepository.Create(basketItem)
	if err != nil {
		return nil, err
	}

	return basketItem, nil
}

func (s *basketItemService) Update(id uint, updateDto *dto.UpdateBasketItemDto) (*entity.BasketItem, error) {
	basketItem, err := s.basketItemRepository.Show(id)
	if err != nil {
		return nil, err
	}

	basketItem.Quantity = updateDto.Quantity

	err = s.basketItemRepository.Update(basketItem)
	if err != nil {
		return nil, err
	}

	return basketItem, nil
}

func (s *basketItemService) Delete(id uint) error {
	basketItem, err := s.basketItemRepository.Show(id)
	if err != nil {
		return err
	}

	return s.basketItemRepository.Delete(basketItem)
}

// Implement the methods of the BasketItemService interface...
