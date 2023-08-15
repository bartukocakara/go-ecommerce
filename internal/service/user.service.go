package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type UserService interface {
	GetUsers(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error)
	GetUserByID(id uint) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	DeleteUser(user *entity.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUsers(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error) {
	return s.userRepository.GetUsers(offset, limit, filter)
}

func (s *userService) GetUserByID(id uint) (*entity.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(user *entity.User) error {
	err := s.userRepository.CreateUser(user)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *userService) UpdateUser(user *entity.User) error {
	err := s.userRepository.UpdateUser(user)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *userService) DeleteUser(user *entity.User) error {
	err := s.userRepository.DeleteUser(user)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}
