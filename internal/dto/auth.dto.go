package dto

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User        UserDTO `json:"user"`
	Role        string  `json:"role"`
	AccessToken string  `json:"access_token"`
}

type RegisterDto struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=2,max=100"`
}

type RegistrationResponse struct {
	User        UserDTO `json:"user"`
	Role        string  `json:"role"`
	AccessToken string  `json:"access_token"`
}

type ForgotPasswordDto struct {
	Email string `json:"email"`
}
