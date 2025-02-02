package routes

import (
	"github.com/gin-gonic/gin"
	"ai-video-repurposer/handlers"
)

// SetupRoutes initializes all API routes
func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "API is running"})
	})

	router.POST("/upload", handlers.UploadVideo)
	router.GET("/access/:encrypted", handlers.AccessVideo)
}
