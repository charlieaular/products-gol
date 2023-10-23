package repositories

import (
	"github.com/charlieaular/products-go/internal/domain/contracts"
	"github.com/charlieaular/products-go/internal/domain/models"
	"gorm.io/gorm"
)

type ProductsRepository struct {
	DB *gorm.DB
}

func NewProductsRepository(db *gorm.DB) contracts.ProductsRepositoryContract {
	return ProductsRepository{DB: db}
}

func (r ProductsRepository) GetAll(filter models.Product) ([]models.Product, error) {
	var products []models.Product

	result := r.DB.Model(models.Product{}).Where("sku like ?", "%"+filter.SKU+"%").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r ProductsRepository) DeleteById(id string) error {
	err := r.DB.Unscoped().Delete(&models.Product{}, id).Error

	return err
}

func (r ProductsRepository) CreateProduct(model models.Product) (models.Product, error) {
	result := r.DB.Model(models.Product{}).Create(&model)

	if result.Error != nil {
		return models.Product{}, nil
	}

	return model, nil
}

func (r ProductsRepository) UpdateProduct(id string, product models.Product) (models.Product, error) {
	result := r.DB.Model(models.Product{}).Where("id = ?", id).Updates(&product)

	if result.Error != nil {
		return models.Product{}, nil
	}

	return product, nil
}
