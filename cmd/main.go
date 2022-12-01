package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"rest-hw"
	"rest-hw/pkg/handler"
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

	handlers := new(handler.Handler)

	server := new(appServer.Server)

	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %v", err.Error())
	}
}
