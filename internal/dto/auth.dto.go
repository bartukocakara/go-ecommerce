package dto

import (
	"ecommerce/internal/entity"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User        entity.User `json:"user"`
	AccessToken string      `json:"access_token"`
}

type RegisterDto struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=2,max=100"`
}

type RegistrationResponse struct {
	User        RegisterDto `json:"user"`
	Role        string      `json:"role"`
	AccessToken string      `json:"access_token"`
}

type ForgetPasswordDto struct {
	Email string `json:"email"`
}
