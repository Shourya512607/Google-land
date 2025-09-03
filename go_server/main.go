package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Use handlers from the same package
	router.GET("/hello", GetHello)
	router.POST("/echo", PostEcho)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Go server listening on port %s", port)
	log.Fatal(router.Run(":" + port))
}
