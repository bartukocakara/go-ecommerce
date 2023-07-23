package dto

import "ecommerce/internal/entity"

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User        entity.User `json:"user"`
	AccessToken string      `json:"access_token"`
}

type RegisterDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegistrationResponse struct {
	User        RegisterDto `json:"user"`
	AccessToken string      `json:"access_token"`
}

type ForgetPasswordDto struct {
	Email string `json:"email"`
}
