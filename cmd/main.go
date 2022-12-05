package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"rest-hw"
	"rest-hw/pkg/handler"
	"rest-hw/pkg/repository"
	"rest-hw/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %v", err.Error())
	}

	mode := viper.GetString("mode")
	port := viper.GetString("port")

	gin.SetMode(mode)

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading environment variables: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PW"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Error initializing Postgres: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(appServer.Server)

	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error running server: %v", err.Error())
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
