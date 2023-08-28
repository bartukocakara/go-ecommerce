package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	List(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error)
	Show(id uint) (*entity.User, error)
	Create(user *dto.CreateUserDTO) (*entity.User, error)
	Update(id uint, user *dto.UpdateUserDTO) error
	Delete(id uint) error
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

func (s *userService) Create(dto *dto.CreateUserDTO) (*entity.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user := &entity.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  string(hashedPassword),
		RoleID:    RoleCustomer,
	}

	user, err := s.userRepository.Create(user)
	if err != nil {
		// Handle error
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(id uint, dto *dto.UpdateUserDTO) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user := &entity.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  string(hashedPassword),
		RoleID:    dto.RoleID,
	}
	err := s.userRepository.Update(id, user)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *userService) Delete(id uint) error {
	err := s.userRepository.Delete(id)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}
