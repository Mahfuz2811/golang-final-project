package repositories

import (
	"final-golang-project/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *models.Product) error {
	// query := "INSERT INTO users (username, email, password, is_verified, verification_token) VALUES(?, ?, ?, ?, ?)"
	// _, err := r.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.IsVerified, user.VerificationToken)
	// if err != nil {
	// 	fmt.Printf("error during user creation: %s", err)
	// }

	// return err

	return r.db.Create(product).Error

}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
