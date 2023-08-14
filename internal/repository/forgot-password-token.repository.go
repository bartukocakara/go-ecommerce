package repository

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type ForgotPasswordTokenRepository interface {
	GetByUserId(userId uint) (*entity.ForgotPasswordToken, error)
	GetByToken(token string) (*entity.ForgotPasswordToken, error)
	CreateToken(Token *entity.ForgotPasswordToken) error
	DeleteToken(Token *entity.ForgotPasswordToken) error
}

type forgotPasswordTokenRepository struct {
	db *gorm.DB
}

func NewForgotPasswordTokenRepository(db *gorm.DB) ForgotPasswordTokenRepository {
	return &forgotPasswordTokenRepository{
		db: db,
	}
}

func (r *forgotPasswordTokenRepository) GetByUserId(userId uint) (*entity.ForgotPasswordToken, error) {
	var forgotPasswordToken entity.ForgotPasswordToken
	result := r.db.Model(&entity.ForgotPasswordToken{}).Where("user_id = ?", userId).First(&forgotPasswordToken)
	if result.Error != nil {
		return nil, result.Error
	}
	return &forgotPasswordToken, nil
}

func (r *forgotPasswordTokenRepository) GetByToken(token string) (*entity.ForgotPasswordToken, error) {
	var ForgotPasswordToken entity.ForgotPasswordToken
	result := r.db.Model(&entity.ForgotPasswordToken{}).Where("token", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ForgotPasswordToken, nil
}

func (r *forgotPasswordTokenRepository) CreateToken(forgotPasswordToken *entity.ForgotPasswordToken) error {
	result := r.db.Create(forgotPasswordToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *forgotPasswordTokenRepository) DeleteToken(Token *entity.ForgotPasswordToken) error {
	result := r.db.Delete(Token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
