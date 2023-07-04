package dto

// UserDTO represents the data transfer object for User entity.
type UserDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// CreateUserDTO represents the data transfer object for creating a new User.
type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UpdateUserDTO represents the data transfer object for updating an existing User.
type UpdateUserDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
