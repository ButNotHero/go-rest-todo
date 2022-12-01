package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"rest-hw"
	"rest-hw/pkg/handler"
	"rest-hw/pkg/repository"
	"rest-hw/pkg/service"
)

func main() {
	PORT := os.Getenv("PORT")
	// debug или release
	ENV := os.Getenv("ENV")

	if PORT == "" {
		PORT = "3000"
	}

	if ENV == "" {
		ENV = "debug"
	}

	gin.SetMode(ENV)

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(appServer.Server)

	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %v", err.Error())
	}
}
