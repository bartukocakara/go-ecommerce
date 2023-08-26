package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type BasketService interface {
	List() ([]entity.Basket, error)
	Show(id uint) (*entity.Basket, error)
	Create(createDto *dto.CreateBasketDto) (*entity.Basket, error)
	Update(id uint, basketDto *dto.UpdateBasketDto) (*entity.Basket, error)
	Delete(id uint) error
}

type basketService struct {
	basketRepository repository.BasketRepository
}

func NewBasketService(basketRepository repository.BasketRepository) BasketService {
	return &basketService{
		basketRepository: basketRepository,
	}
}

func (s *basketService) List() ([]entity.Basket, error) {
	return s.basketRepository.List()
}

func (s *basketService) Show(id uint) (*entity.Basket, error) {
	return s.basketRepository.Show(id)
}

func (s *basketService) Create(createDto *dto.CreateBasketDto) (*entity.Basket, error) {
	basket := &entity.Basket{
		TotalPrice: createDto.TotalPrice,
	}

	err := s.basketRepository.Create(basket)
	if err != nil {
		return nil, err
	}

	return basket, nil
}

func (s *basketService) Update(id uint, updateDto *dto.UpdateBasketDto) (*entity.Basket, error) {
	Basket, err := s.basketRepository.Show(id)
	if err != nil {
		return nil, err
	}

	Basket.TotalPrice = updateDto.TotalPrice

	err = s.basketRepository.Update(Basket)
	if err != nil {
		return nil, err
	}

	return Basket, nil
}

func (s *basketService) Delete(id uint) error {
	Basket, err := s.basketRepository.Show(id)
	if err != nil {
		return err
	}

	return s.basketRepository.Delete(Basket)
}

// Implement the methods of the BasketService interface...
