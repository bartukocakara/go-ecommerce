package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type BasketService interface {
	GetBaskets() ([]entity.Basket, error)
	GetBasketByID(id uint) (*entity.Basket, error)
	CreateBasket(createDto *dto.CreateBasketDto) (*entity.Basket, error)
	UpdateBasket(id uint, basketDto *dto.UpdateBasketDto) (*entity.Basket, error)
	DeleteBasket(id uint) error
}

type basketService struct {
	basketRepository repository.BasketRepository
}

func NewBasketService(basketRepository repository.BasketRepository) BasketService {
	return &basketService{
		basketRepository: basketRepository,
	}
}

func (s *basketService) GetBaskets() ([]entity.Basket, error) {
	return s.basketRepository.GetBaskets()
}

func (s *basketService) GetBasketByID(id uint) (*entity.Basket, error) {
	return s.basketRepository.GetBasketByID(id)
}

func (s *basketService) CreateBasket(createDto *dto.CreateBasketDto) (*entity.Basket, error) {
	basket := &entity.Basket{
		TotalPrice: createDto.TotalPrice,
	}

	err := s.basketRepository.CreateBasket(basket)
	if err != nil {
		return nil, err
	}

	return basket, nil
}

func (s *basketService) UpdateBasket(id uint, updateDto *dto.UpdateBasketDto) (*entity.Basket, error) {
	Basket, err := s.basketRepository.GetBasketByID(id)
	if err != nil {
		return nil, err
	}

	Basket.TotalPrice = updateDto.TotalPrice

	err = s.basketRepository.UpdateBasket(Basket)
	if err != nil {
		return nil, err
	}

	return Basket, nil
}

func (s *basketService) DeleteBasket(id uint) error {
	Basket, err := s.basketRepository.GetBasketByID(id)
	if err != nil {
		return err
	}

	return s.basketRepository.DeleteBasket(Basket)
}

// Implement the methods of the BasketService interface...
