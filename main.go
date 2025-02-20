package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nomankhokhar/GoTonyUrl/api/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()

	setupRoutes(router)

	PORT := os.Getenv("APP_PORT")

	if PORT == "" {
		PORT = "8000"
	}

	log.Fatal(router.Run(":" + PORT))
}

func setupRoutes(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.DELETE("/api/v1/:shortID", routes.DeleteURL)
	router.PUT("/api/v1/:shortID", routes.EditURL)
}
