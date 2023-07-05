package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Register(registerDto dto.RegisterDto) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user entity
	user := &entity.User{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Email:     registerDto.Email,
		Password:  string(hashedPassword),
	}

	// Save the user in the repository
	err = s.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(loginDto dto.LoginDto) (string, error) {
	// Retrieve the user by email
	user, err := s.UserRepository.GetUserByEmail(loginDto.Email)
	if err != nil {
		return "", err
	}

	// Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		return "", err
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		// Add more claims as needed
	})

	// Sign the token with a secret key
	// Replace "YOUR_SECRET_KEY" with your actual secret key
	tokenString, err := token.SignedString([]byte("YOUR_SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
