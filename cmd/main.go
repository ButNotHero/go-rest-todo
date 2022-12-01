package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"rest-hw"
	"rest-hw/pkg/handler"
	"rest-hw/pkg/repository"
	"rest-hw/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %v", err.Error())
	}

	env := viper.GetString("env")
	port := viper.GetString("port")

	gin.SetMode(env)

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(appServer.Server)

	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %v", err.Error())
	}
}

func initConfig() error {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
