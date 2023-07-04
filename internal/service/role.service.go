package service

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type RoleService interface {
	GetRoles() ([]entity.Role, error)
	GetRoleByID(id uint) (*entity.Role, error)
	CreateRole(Role *entity.Role) error
	UpdateRole(Role *entity.Role) error
	DeleteRole(Role *entity.Role) error
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepository,
	}
}

func (s *roleService) GetRoles() ([]entity.Role, error) {
	roles, err := s.roleRepository.GetRoles()
	if err != nil {
		// Handle error
		return nil, err
	}

	return roles, nil
}

func (s *roleService) GetRoleByID(id uint) (*entity.Role, error) {
	role, err := s.roleRepository.GetRoleByID(id)
	if err != nil {
		// Handle error
		return nil, err
	}

	return role, nil
}

func (s *roleService) CreateRole(role *entity.Role) error {
	err := s.roleRepository.CreateRole(role)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *roleService) UpdateRole(role *entity.Role) error {
	err := s.roleRepository.UpdateRole(role)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (s *roleService) DeleteRole(role *entity.Role) error {
	err := s.roleRepository.DeleteRole(role)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}
