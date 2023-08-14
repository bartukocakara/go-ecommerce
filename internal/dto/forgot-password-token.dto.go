package dto

import "time"

// CreateUserDTO represents the data transfer object for creating a new User.
type CreateForgotPasswordDTO struct {
	UserID    uint   `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt time.Time
}
