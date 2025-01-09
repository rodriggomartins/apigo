package main

import (
	"fmt"
	"log"

	"api/src/configuration/logger"
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
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":9090"); err != nil {
		log.Fatal(err)
	}

}
