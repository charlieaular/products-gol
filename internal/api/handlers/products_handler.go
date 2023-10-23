package handlers

import (
	"net/http"

	"github.com/charlieaular/products-go/internal/domain/contracts"
	dtos "github.com/charlieaular/products-go/internal/domain/dtos/products"
	"github.com/charlieaular/products-go/internal/domain/models"
	"github.com/charlieaular/products-go/internal/shared"
	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	service contracts.ProductsServiceContract
}

func NewProductsHandler(service contracts.ProductsServiceContract) ProductsHandler {
	return ProductsHandler{service: service}
}

func (h ProductsHandler) GetAll(c *gin.Context) {
	var filter models.Product

	err := c.ShouldBindQuery(&filter)
	if shared.HandleError(c, err) {
		return
	}

	products, err := h.service.GetAll(filter)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})
}

func (h ProductsHandler) DeleteProduct(c *gin.Context) {
	productId := c.Param("id")

	err := h.service.DeleteById(productId)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func (h ProductsHandler) CreateProduct(c *gin.Context) {
	var createProductDto dtos.CreateProductDTO

	err := c.ShouldBindJSON(&createProductDto)

	if shared.HandleError(c, err) {
		return
	}

	model := models.Product{
		SKU:   createProductDto.SKU,
		Name:  createProductDto.Name,
		Brand: createProductDto.Brand,
		Size:  createProductDto.Size,
		Price: createProductDto.Price,
	}

	result, err := h.service.CreateProduct(model)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"product": result,
	})
}

func (h ProductsHandler) UpdateProduct(c *gin.Context) {
	var updateProductDto dtos.UpdateProductDTO

	err := c.ShouldBindJSON(&updateProductDto)

	if shared.HandleError(c, err) {
		return
	}

	model := models.Product{
		SKU:   updateProductDto.SKU,
		Name:  updateProductDto.Name,
		Brand: updateProductDto.Brand,
		Size:  updateProductDto.Size,
		Price: updateProductDto.Price,
	}

	productId := c.Param("id")

	result, err := h.service.UpdateProduct(productId, model)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": result,
	})
}
