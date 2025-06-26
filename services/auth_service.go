package services

import (
	"errors"
	"final-golang-project/models"
	"final-golang-project/repositories"
	utils "final-golang-project/utils"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo    repositories.UserRepository
	emailSender EmailSender
}

// concreate type
// no way mock

func NewAuthServe(userRepo repositories.UserRepository, emailSender EmailSender) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		emailSender: emailSender,
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
	// if existingUser != nil || err != nil {
	// 	return fmt.Errorf("user already exists or error: %s", err)
	// }

	if err != nil {
		return fmt.Errorf("failed to check existing user: %s", err)
	}

	if existingUser != nil {
		return fmt.Errorf("user already exists")
	}

	if err := s.userRepo.Create(user); err != nil {
		return err
	}

	// send verification email
	if err := s.emailSender.SendVerificationEmail(email, verificationToken); err != nil {
		return fmt.Errorf("failed to send verification email: %s", err)
	}

	return nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}

	// user->password, provided password

	error := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if error != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}
