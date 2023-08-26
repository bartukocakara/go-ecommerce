package repository

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	List(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error)
	Show(id uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserRoleNameByID(userID uint) (string, error)
	GetPermissionsByUserID(userID uint) ([]string, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) error
	Delete(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) List(offset, limit int, filter *dto.FilterUserDTO) ([]*entity.User, int, error) {
	var users []*entity.User

	db := r.db.Model(&entity.User{})

	if filter != nil {
		if filter.FirstName != "" {
			db = db.Where("first_name ILIKE ?", "%"+filter.FirstName+"%")
		}
		if filter.LastName != "" {
			db = db.Where("last_name ILIKE ?", "%"+filter.LastName+"%")
		}
		if filter.Email != "" {
			db = db.Where("email ILIKE ?", "%"+filter.Email+"%")
		}
	}
	count, err := CountTotal(db, users)
	if err != nil {
		return nil, 0, err
	}
	db = Paginate(db, offset, limit)

	if err := db.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(count), nil
}

func (r *userRepository) Show(id uint) (*entity.User, error) {
	var user entity.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := r.db.Preload("Role").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserRoleNameByID(userID uint) (string, error) {
	var roleName string
	if err := r.db.Model(&entity.User{}).
		Select("roles.name").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("users.id = ?", userID).Scan(&roleName).Error; err != nil {
		return "", err
	}
	return roleName, nil
}

func (r *userRepository) GetPermissionsByUserID(userID uint) ([]string, error) {
	var permissions []string
	if err := r.db.Model(&entity.User{}).
		Select("permissions.name").
		Joins("JOIN roles ON users.role_id = roles.id").
		Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("users.id = ?", userID).
		Pluck("permissions.name", &permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	// If the role ID is provided in the user entity, fetch the role from the database
	if user.RoleID != 0 {
		var role entity.Role
		if err := r.db.First(&role, user.RoleID).Error; err != nil {
			return nil, err
		}
		// Set the role association on the user entity
		user.Role = role
	}

	// Create the user record
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) Delete(user *entity.User) error {
	result := r.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
