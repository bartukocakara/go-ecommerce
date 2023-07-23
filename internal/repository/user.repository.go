package repository

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(offset, limit int) ([]*entity.User, int, error)
	GetUserByID(id uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	DeleteUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers(offset, limit int) ([]*entity.User, int, error) {
	var users []*entity.User
	var total int64

	// Count total number of users
	if err := r.db.Model(&entity.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch users with pagination
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *userRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *entity.User) error {
    result := r.db.Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *userRepository) UpdateUser(user *entity.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) DeleteUser(user *entity.User) error {
	result := r.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
