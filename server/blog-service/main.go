package main

import (
	"os"

	"github.com/KwesiLarbi/blog-service/middleware"
	"github.com/KwesiLarbi/blog-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	
	// User routes
	routes.UserRoutes(router)

	// All other routes requiring authentication are after the auth middleware
	router.Use(middleware.Authentication())
	
	// Dummy API-1
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// Dummy API-2
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}