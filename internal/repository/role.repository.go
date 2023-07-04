package repository

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoles() ([]entity.Role, error)
	GetRoleByID(id uint) (*entity.Role, error)
	CreateRole(Role *entity.Role) error
	UpdateRole(Role *entity.Role) error
	DeleteRole(Role *entity.Role) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetRoles() ([]entity.Role, error) {
	var Roles []entity.Role
	result := r.db.Find(&Roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return Roles, nil
}

func (r *roleRepository) GetRoleByID(id uint) (*entity.Role, error) {
	var Role entity.Role
	result := r.db.First(&Role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Role, nil
}

func (r *roleRepository) CreateRole(Role *entity.Role) error {
	result := r.db.Create(Role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *roleRepository) UpdateRole(Role *entity.Role) error {
	result := r.db.Save(Role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *roleRepository) DeleteRole(Role *entity.Role) error {
	result := r.db.Delete(Role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
