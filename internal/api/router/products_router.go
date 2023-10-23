package router

import (
	"github.com/charlieaular/products-go/internal/api/handlers"
	"github.com/charlieaular/products-go/internal/domain/services"
	"github.com/charlieaular/products-go/internal/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductsRouter(router *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewProductsRepository(db)
	service := services.NewProductsService(repo)
	handler := handlers.NewProductsHandler(service)

	routes := router.Group("products")
	{
		routes.GET("", handler.GetAll)
		routes.POST("", handler.CreateProduct)
		routes.PUT(":id", handler.UpdateProduct)
		routes.DELETE(":id", handler.DeleteProduct)
	}
}
