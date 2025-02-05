package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"ai-video-repurposer/routes"
	"github.com/gin-contrib/cors"
    "gorm.io/gorm"
	"ai-video-repurposer/initializers"
)
var DB *gorm.DB

func main() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	// Initialize Gin
	router := gin.Default()
	router.Use(cors.Default())


	// Set up routes
	routes.SetupRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("🚀 Server running on http://localhost:" + port)
	router.Run(":" + port)
}
