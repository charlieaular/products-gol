package repositories_test

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/charlieaular/products-go/internal/domain/contracts"
	"github.com/charlieaular/products-go/internal/domain/models"
	"github.com/charlieaular/products-go/internal/infrastructure/database"
	"github.com/charlieaular/products-go/internal/infrastructure/repositories"
)

type ProductsRepositorySuite struct {
	suite.Suite
	DB   *gorm.DB
	repo contracts.ProductsRepositoryContract
}

func (suite *ProductsRepositorySuite) SetupTest() {
	viper.SetConfigFile("../../../.env.testing")
	viper.SetConfigType("dotenv")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	mysqlIns := database.MysqlConn

	suite.DB = mysqlIns.Init()
	suite.repo = repositories.NewProductsRepository(suite.DB)
}

func (suite *ProductsRepositorySuite) TearDownTest() {
	suite.DB.Exec("TRUNCATE TABLE products")
}

func (suite *ProductsRepositorySuite) TestGetAll() {
	products, err := suite.repo.GetAll(models.Product{})
	assert.NoError(suite.T(), err, "get all products fails")
	assert.Equal(suite.T(), 0, len(products), "same length of products from db")

}

func (suite *ProductsRepositorySuite) TestCreate() {

	model := models.Product{
		SKU:   "SKU",
		Name:  "NAME",
		Brand: "BRAND",
		Size:  "SIZE",
		Price: 1000,
	}

	product, err := suite.repo.CreateProduct(model)
	assert.NoError(suite.T(), err, "product creation fails")
	assert.Equal(suite.T(), product.SKU, model.SKU, "same product returned from db")

}

func (suite *ProductsRepositorySuite) TestUpdate() {
	model := models.Product{
		SKU:   "SKU",
		Name:  "NAME",
		Brand: "BRAND",
		Size:  "SIZE",
		Price: 1000,
	}

	product, _ := suite.repo.CreateProduct(model)

	model.Name = "THIS IS NEW"

	product, err := suite.repo.UpdateProduct(fmt.Sprintf("%v", model.ID), model)
	assert.NoError(suite.T(), err, "user creation fails")
	assert.Equal(suite.T(), product.Name, model.Name, "same updated product returned from db")

}

func (suite *ProductsRepositorySuite) TestGetAllFiltered() {
	model := models.Product{
		SKU:   "SKU-1",
		Name:  "NAME",
		Brand: "BRAND",
		Size:  "SIZE",
		Price: 1000,
	}

	suite.repo.CreateProduct(model)

	model.SKU = "SKU-2"
	suite.repo.CreateProduct(model)

	products, err := suite.repo.GetAll(models.Product{SKU: "SKU-2"})
	assert.NoError(suite.T(), err, "get all products filtered fails")
	assert.Equal(suite.T(), 1, len(products), "same length of products filtered from db")

}

func (suite *ProductsRepositorySuite) TestDeleteById() {
	model := models.Product{
		SKU:   "SKU-1",
		Name:  "NAME",
		Brand: "BRAND",
		Size:  "SIZE",
		Price: 1000,
	}

	suite.repo.CreateProduct(model)

	model.SKU = "SKU-2"
	productToDelete, _ := suite.repo.CreateProduct(model)

	model.SKU = "SKU-3"
	suite.repo.CreateProduct(model)

	err := suite.repo.DeleteById(fmt.Sprintf("%v", productToDelete.ID))
	assert.NoError(suite.T(), err, "user deletion fails")

	products, _ := suite.repo.GetAll(models.Product{})

	assert.Equal(suite.T(), 2, len(products), "same length of products from db")

}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductsRepositorySuite))
}
