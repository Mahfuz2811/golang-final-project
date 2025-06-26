package services

import (
	"final-golang-project/models"
	"final-golang-project/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) ListByUserEmail(userEmail string) ([]models.Product, error) {
	products, err := s.repo.ListByUserEmail(userEmail)
	if err != nil {
		return nil, err
	}
	return products, nil
}
