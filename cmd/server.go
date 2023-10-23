package server

import (
	"github.com/charlieaular/products-go/internal/api/router"
	"github.com/charlieaular/products-go/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SetupDBInstance() *gorm.DB {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	mysqlIns := database.MysqlConn

	db := mysqlIns.Init()
	return db
}

func SetUpRouter(db *gorm.DB) *gin.Engine {

	engine := gin.Default()

	v1 := engine.Group("v1")

	router.NewProductsRouter(v1, db)

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Route not found"})
	})

	return engine

}
