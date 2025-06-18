package services

import (
	"final-golang-project/models"
	"final-golang-project/repositories"
	utils "final-golang-project/utlis"
	"fmt"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

// concreate type
// no way mock

func NewAuthServe(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) RegisterUser(username, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	verificationToken := uuid.New().String()

	user := models.User{
		Username:          username,
		Email:             email,
		PasswordHash:      hashedPassword,
		IsVerified:        false,
		VerificationToken: verificationToken,
	}

	existingUser, err := s.userRepo.GetByEmail(email)
	fmt.Println("Existing User:", existingUser)
	if existingUser != nil || err != nil {
		return fmt.Errorf("user already exists or error: %s", err)
	}

	if err := s.userRepo.Create(user); err != nil {
		return err
	}

	// send verification email
	utils.SendVerificationEmail(email, verificationToken)

	return nil
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}
