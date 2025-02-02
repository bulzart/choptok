package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/pusher/pusher-http-go"
)

// Storage for temporary files
var fileStorage = struct {
	sync.Mutex
	files map[string]string
}{files: make(map[string]string)}

var pusherClient = pusher.Client{
    AppID:   "1602604",
    Key:     "1ac366cc7e1e8ef0dd7c",
    Secret:  "0e1eb7be834e236d6949",
    Cluster: "eu",
    Secure:  true,
}

// UploadVideo handles file uploads
func UploadVideo(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}
	defer file.Close()

	// Generate random encrypted filename
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	encryptedName := hex.EncodeToString(randomBytes) + filepath.Ext(header.Filename)

	// Save file
	savePath := "uploads/" + encryptedName
	out, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer out.Close()
	io.Copy(out, file)

	// Store in memory
	fileStorage.Lock()
	fileStorage.files[encryptedName] = savePath
	fileStorage.Unlock()

	// Auto-delete after 10 minutes
	go func(filename string) {
		time.Sleep(10 * time.Minute)
		fileStorage.Lock()
		delete(fileStorage.files, filename)
		fileStorage.Unlock()
		os.Remove("uploads/" + filename)
	}(encryptedName)
	addSubtitlesToVideo(encryptedName)
	pusherClient.Trigger("my-channel", "my-event", map[string]string{"message": "File uploaded successfully"})
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded", "access_url": "/access/" + encryptedName})
}

// AccessVideo serves the uploaded video via encrypted link
func AccessVideo(c *gin.Context) {
	encryptedName := c.Param("encrypted")

	fileStorage.Lock()
	filePath, exists := fileStorage.files[encryptedName]
	fileStorage.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "File expired or not found"})
		return
	}

	c.File(filePath)
}
