package dto

// RoleDTO represents the data transfer object for Role entity.
type RoleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CreateRoleDTO represents the data transfer object for creating a new Role.
type CreateRoleDTO struct {
	Name string `json:"name"`
}

// UpdateRoleDTO represents the data transfer object for updating an existing Role.
type UpdateRoleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
