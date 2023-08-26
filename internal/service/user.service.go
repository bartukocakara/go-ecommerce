package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type UserService interface {
	List(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error)
	Show(id uint) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) List(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error) {
	return s.userRepository.List(offset, limit, filter)
}

func (s *userService) Show(id uint) (*entity.User, error) {
	user, err := s.userRepository.Show(id)
	if err != nil {
		// Handle error
		return nil, err
	}

	return user, nil
}

func (s *userService) Create(user *entity.User) error {
	err := s.userRepository.Create(user)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *userService) Update(user *entity.User) error {
	err := s.userRepository.Update(user)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *userService) Delete(user *entity.User) error {
	err := s.userRepository.Delete(user)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}
