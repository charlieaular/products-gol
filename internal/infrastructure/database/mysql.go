package database

import (
	"fmt"

	"github.com/charlieaular/products-go/internal/domain/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct{}

var MysqlConn Mysql

func (p *Mysql) Init() *gorm.DB {
	dbHost := viper.Get("DB_HOST")
	dbUser := viper.Get("DB_USER")
	dbPassword := viper.Get("DB_PASSWORD")
	dbPort := viper.Get("DB_PORT")
	dbName := viper.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Product{},
	)

	return db

}
