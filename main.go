package main

import (
	"fmt"

	server "github.com/charlieaular/products-go/cmd"
	"github.com/spf13/viper"
)

func main() {

	db := server.SetupDBInstance()

	engine := server.SetUpRouter(db)

	port := viper.Get("PORT")
	portFormatted := fmt.Sprintf(":%s", port)

	engine.Run(portFormatted)
}
