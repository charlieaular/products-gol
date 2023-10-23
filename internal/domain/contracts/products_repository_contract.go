package contracts

import "github.com/charlieaular/products-go/internal/domain/models"

type ProductsRepositoryContract interface {
	GetAll(models.Product) ([]models.Product, error)
	DeleteById(id string) error
	CreateProduct(models.Product) (models.Product, error)
	UpdateProduct(string, models.Product) (models.Product, error)
}
