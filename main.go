package main

import (
	"fmt"
	"log"

	"api/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
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
