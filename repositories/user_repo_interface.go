package repositories

import "final-golang-project/models"

type UserRepository interface {
	Create(user models.User) error
	GetByEmail(email string) (*models.User, error)
}
