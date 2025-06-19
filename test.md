package services

import (
"errors"
"final-golang-project/models"
"testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "golang.org/x/crypto/bcrypt"

)

type MockUserRepo struct {
mock.Mock
}

func (m \*MockUserRepo) Create(user models.User) error {
args := m.Called(user)
return args.Error(0)
}

func (m *MockUserRepo) GetByEmail(email string) (*models.User, error) {
args := m.Called(email)
return args.Get(0).(\*models.User), args.Error(1)
}

type MockEmailSender struct {
mock.Mock
}

func (m \*MockEmailSender) SendVerificationEmail(email, token string) {
m.Called(email, token)
}

func TestRegisterUser_Success(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    mockRepo.On("GetByEmail", "test@example.com").Return((*models.User)(nil), nil)
    mockRepo.On("Create", mock.AnythingOfType("models.User")).Return(nil)
    mockEmail.On("SendVerificationEmail", "test@example.com", mock.AnythingOfType("string")).Return()

    err := service.RegisterUser("testuser", "test@example.com", "password123")
    assert.NoError(t, err)

    mockRepo.AssertExpectations(t)
    mockEmail.AssertExpectations(t)

}

func TestRegisterUser_UserAlreadyExists(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    existingUser := &models.User{Email: "test@example.com"}
    mockRepo.On("GetByEmail", "test@example.com").Return(existingUser, nil)

    err := service.RegisterUser("testuser", "test@example.com", "password123")
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "user already exists")

}

func TestRegisterUser_GetByEmailError(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    mockRepo.On("GetByEmail", "test@example.com").Return((*models.User)(nil), errors.New("db error"))

    err := service.RegisterUser("testuser", "test@example.com", "password123")
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "failed to check existing user")

}

func TestLogin_Success(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
    user := &models.User{
    	Email:        "test@example.com",
    	PasswordHash: string(hashed),
    }

    mockRepo.On("GetByEmail", "test@example.com").Return(user, nil)

    result, err := service.Login("test@example.com", "password123")
    assert.NoError(t, err)
    assert.Equal(t, user, result)

}

func TestLogin_InvalidPassword(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
    user := &models.User{
    	Email:        "test@example.com",
    	PasswordHash: string(hashed),
    }

    mockRepo.On("GetByEmail", "test@example.com").Return(user, nil)

    result, err := service.Login("test@example.com", "wrongpassword")
    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Equal(t, "invalid email or password", err.Error())

}

func TestLogin_UserNotFound(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    mockRepo.On("GetByEmail", "missing@example.com").Return((*models.User)(nil), nil)

    result, err := service.Login("missing@example.com", "password123")
    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Equal(t, "invalid email or password", err.Error())

}

func TestGetUserByEmail_Success(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    expected := &models.User{Email: "found@example.com"}
    mockRepo.On("GetByEmail", "found@example.com").Return(expected, nil)

    user, err := service.GetUserByEmail("found@example.com")
    assert.NoError(t, err)
    assert.Equal(t, expected, user)

}

func TestGetUserByEmail_Failure(t \*testing.T) {
mockRepo := new(MockUserRepo)
mockEmail := new(MockEmailSender)
service := NewAuthServe(mockRepo, mockEmail)

    mockRepo.On("GetByEmail", "error@example.com").Return((*models.User)(nil), errors.New("db error"))

    user, err := service.GetUserByEmail("error@example.com")
    assert.Error(t, err)
    assert.Nil(t, user)

}
