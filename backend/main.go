package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

func Shorten(url string) string {
	// Implement your URL shortening logic here
	// For example, you can use a hash function to generate a short code
	// and store the mapping in a database or in-memory data structure
	

	return "shortened_url" // Placeholder return value
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a simple route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":" + os.Getenv("PORT"))
	log.Println("Server is running on port " + os.Getenv("PORT"))
}