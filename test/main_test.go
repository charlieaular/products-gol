package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/charlieaular/products-go/cmd"
	dtos "github.com/charlieaular/products-go/internal/domain/dtos/products"
	"github.com/charlieaular/products-go/internal/domain/models"
	"github.com/charlieaular/products-go/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServerSuite struct {
	suite.Suite
	server *gin.Engine
	DB     *gorm.DB
}

func (suite *ServerSuite) SetupTest() {
	viper.SetConfigFile("../.env.testing")
	viper.SetConfigType("dotenv")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	mysqlIns := database.MysqlConn

	db := mysqlIns.Init()
	gin.SetMode(gin.ReleaseMode)

	suite.DB = db
	suite.server = server.SetUpRouter(db)

}

func (suite *ServerSuite) TearDownTest() {
	suite.DB.Exec("TRUNCATE TABLE products")
}

func (suite *ServerSuite) TestGetAll() {

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/products", nil)
	suite.server.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *ServerSuite) TestGetAllFiltered() {

	model := models.Product{
		SKU:   "filter",
		Name:  "test",
		Brand: "test",
		Size:  "test",
		Price: 100,
	}

	suite.DB.Create(&model)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/products?sku=filter", nil)
	suite.server.ServeHTTP(w, req)

	var response struct {
		Status   bool
		Products []models.Product
	}

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Equal(suite.T(), 1, len(response.Products))
}

func (suite *ServerSuite) TestCreateProduct() {

	request := dtos.CreateProductDTO{
		SKU:   "SKU-1",
		Name:  "Product Name",
		Brand: "Product Brand",
		Price: 2000,
		Size:  "S",
	}

	jsonValue, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/v1/products", bytes.NewBuffer(jsonValue))
	suite.server.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
}

func (suite *ServerSuite) TestUpdateProduct() {

	requestCreate := dtos.CreateProductDTO{
		SKU:   "SKU-2",
		Name:  "New Product Name",
		Brand: "New Product Brand",
		Price: 1500,
		Size:  "M",
	}

	request := dtos.UpdateProductDTO{
		CreateProductDTO: requestCreate,
	}

	jsonValue, _ := json.Marshal(request)

	model := models.Product{
		SKU:   "test",
		Name:  "test",
		Brand: "test",
		Size:  "test",
		Price: 100,
	}

	suite.DB.Create(&model)

	var lastProduct models.Product

	suite.DB.Model(models.Product{}).Order("created_at desc").First(&lastProduct)

	route := fmt.Sprintf("/v1/products/%v", lastProduct.ID)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", route, bytes.NewBuffer(jsonValue))
	suite.server.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *ServerSuite) TestDeleteProduct() {
	var lastProduct models.Product

	model := models.Product{
		SKU:   "test",
		Name:  "test",
		Brand: "test",
		Size:  "test",
		Price: 100,
	}

	suite.DB.Create(&model)

	suite.DB.Model(models.Product{}).Order("created_at desc").First(&lastProduct)
	route := fmt.Sprintf("/v1/products/%v", lastProduct.ID)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", route, nil)
	suite.server.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
