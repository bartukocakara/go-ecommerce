package dto

// RoleDTO represents the data transfer object for Role entity.
type RoleDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// CreateRoleDTO represents the data transfer object for creating a new Role.
type CreateRoleDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UpdateRoleDTO represents the data transfer object for updating an existing Role.
type UpdateRoleDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
