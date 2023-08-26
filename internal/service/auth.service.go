package service

import (
	"crypto/rand"
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/enums"
	"ecommerce/internal/repository"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	RoleAdmin    = 1
	RoleCustomer = 2
)

type AuthService struct {
	UserRepository                repository.UserRepository
	ForgotPasswordTokenRepository repository.ForgotPasswordTokenRepository
	// MailService    MailService
}

func NewAuthService(userRepository repository.UserRepository,
	forgotPasswordTokenRepository repository.ForgotPasswordTokenRepository) *AuthService {
	return &AuthService{
		UserRepository:                userRepository,
		ForgotPasswordTokenRepository: forgotPasswordTokenRepository,
	}
}

func (s *AuthService) Register(registerDto dto.RegisterDto) (*dto.RegistrationResponse, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// Create a new user entity
	user := &entity.User{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Email:     registerDto.Email,
		RoleID:    RoleCustomer,
		Password:  string(hashedPassword),
	}
	// Save the user in the repository
	user, err = s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   registerDto.Email,
		"user_id": user.ID,
		"role":    user.Role.Name,
	})

	// Sign the token with a secret key (replace "your-secret-key" with your actual secret key)
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return nil, err
	}

	response := &dto.RegistrationResponse{
		User: dto.UserDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email},
		Role:        user.Role.Name,
		AccessToken: tokenString,
	}

	return response, nil
}

func (s *AuthService) Login(loginDto dto.LoginDto) (*dto.LoginResponse, error) {
	// Retrieve the user by email
	user, err := s.UserRepository.GetUserByEmail(loginDto.Email)
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		return nil, err
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"user_id": user.ID,
		"role":    user.Role.Name,
		// Add more claims as needed
	})

	// Sign the token with a secret key
	// Replace "YOUR_SECRET_KEY" with your actual secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		User: dto.UserDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email},
		Role: user.Role.Name,

		AccessToken: tokenString,
	}

	return response, nil
}

func (s *AuthService) ForgotPassword(forgotPasswordDto dto.ForgotPasswordDto) (string, enums.ForgotPasswordResult, error) {
	// Check if the user exists in the database
	user, err := s.UserRepository.GetUserByEmail(forgotPasswordDto.Email)
	if err != nil {
		return "User not found", enums.UserNotFound, err
	}
	fmt.Print(user)
	// Get forgot password by user id
	tokenFound, err := s.ForgotPasswordTokenRepository.GetByUserId(user.ID)
	if tokenFound != nil {
		return "Token already exist. Expiring time is 1 day", enums.TokenAlreadyExists, nil
	}
	// Generate a password reset token for the user
	resetToken, err := GeneratePasswordResetToken(user.ID)
	if err != nil {
		return "", enums.TokenCouldntGenerated, nil
	}
	expiresAt := time.Now().Add(24 * time.Hour)

	var forgotPasswordData = entity.ForgotPasswordToken{
		UserID:    user.ID,
		Token:     resetToken,
		ExpiresAt: expiresAt, // Set the calculated expiration time
	}

	data := s.ForgotPasswordTokenRepository.CreateToken(&forgotPasswordData)
	if data != nil {
		return "Token created successfully", enums.Success, nil
	}

	// Send the password reset email to the user
	// err = s.MailService.SendForgetPasswordEmail(user.Email, resetToken)
	// if err != nil {
	// 	return err
	// }

	return resetToken, enums.Success, nil
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
