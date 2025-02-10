package main

import (
	"fmt"
	"log"

	"api/src/configuration/database/mongodb"
	"api/src/configuration/logger"
	"api/src/controller"
	"api/src/model/service"
	"api/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Sobre a aplicacao")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error:", err)
	}

	mongodb.InitConnection()

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":9090"); err != nil {
		log.Fatal(err)
	}

}
