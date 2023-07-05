package service

import (
	"crypto/rand"
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository repository.UserRepository
	// MailService    MailService
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

func (s *AuthService) ForgetPassword(forgetPasswordDto dto.ForgetPasswordDto) (string, error) {
	// Check if the user exists in the database
	user, err := s.UserRepository.GetUserByEmail(forgetPasswordDto.Email)
	if err != nil {
		return "", err
	}

	// Generate a password reset token for the user
	resetToken, err := GeneratePasswordResetToken(user.ID)
	if err != nil {
		return "", err
	}

	// Send the password reset email to the user
	// err = s.MailService.SendForgetPasswordEmail(user.Email, resetToken)
	// if err != nil {
	// 	return err
	// }

	return resetToken, nil
}

func GeneratePasswordResetToken(userID uint) (string, error) {
	// Generate a random token
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	// Encode the token to base64 string
	tokenString := base64.URLEncoding.EncodeToString(token)

	// Generate a unique reset token using user ID and timestamp
	resetToken := fmt.Sprintf("%d_%s_%d", userID, tokenString, time.Now().Unix())

	return resetToken, nil
}
