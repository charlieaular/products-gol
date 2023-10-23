package services

import (
	"github.com/charlieaular/products-go/internal/domain/contracts"
	"github.com/charlieaular/products-go/internal/domain/models"
)

type ProductsService struct {
	repo contracts.ProductsRepositoryContract
}

func NewProductsService(repo contracts.ProductsRepositoryContract) contracts.ProductsServiceContract {
	return ProductsService{repo: repo}
}

func (s ProductsService) GetAll(filter models.Product) ([]models.Product, error) {
	return s.repo.GetAll(filter)
}

func (s ProductsService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s ProductsService) CreateProduct(model models.Product) (models.Product, error) {
	return s.repo.CreateProduct(model)
}

func (s ProductsService) UpdateProduct(id string, model models.Product) (models.Product, error) {
	return s.repo.UpdateProduct(id, model)
}
