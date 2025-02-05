package initializers

import "ai-video-repurposer/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}