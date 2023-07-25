package repository

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error)
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

func (r *userRepository) GetUsers(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error) {

	var users []*entity.User
	var total int64

	db := r.db.Model(&entity.User{})

	// Apply filters based on provided parameters
	if filter != nil {
		if filter.FirstName != "" {
			db = db.Where("first_name LIKE ?", "%"+filter.FirstName+"%")
		}
		if filter.LastName != "" {
			db = db.Where("last_name LIKE ?", "%"+filter.LastName+"%")
		}
		if filter.Email != "" {
			db = db.Where("email LIKE ?", "%"+filter.Email+"%")
		}
	}

	// Count total number of users based on applied filters
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch users with pagination and applied filters
	if err := db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	fmt.Print(users)

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
