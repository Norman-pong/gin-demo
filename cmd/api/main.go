package main

import (
	"log"

	"zhiming.cool/go/internal/routes"
	"zhiming.cool/go/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}
}

func main() {
	initConfig()
	port := viper.GetString("server.port")
	r := gin.Default()

	userService := services.NewUserService("")
	routes.RegisterRoutes(r, userService)
	r.Run(":" + port)
}
