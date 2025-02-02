package handlers

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"context"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/sashabaranov/go-openai"
	"os"
	"crypto/rand"
	"encoding/hex"
)

// ProcessVideo extracts key moments and splits clips
func ProcessVideo(c *gin.Context,encryptedName string){
	encryptedName = encryptedName

	// Check if file exists
	filePath, exists := fileStorage.files[encryptedName]
	if !exists {
		c.JSON(404, gin.H{"error": "File expired or not found"})
		return
	}

	// Extract key moments using OpenAI
	keyMoments, err := extractKeyMomentsFromTranscript(filePath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to extract key moments"})
		return
	}

	// Split video into clips
	clipPaths, err := splitVideo(filePath, keyMoments)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to split video"})
		return
	}

	c.JSON(200, gin.H{"message": "Processing complete", "clips": clipPaths})
}

// Extracts key moments using OpenAI
func extractKeyMomentsFromTranscript(filePath string) ([]string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(context.Background(),
	openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Extract 3 key moments (timestamps) from this video for TikTok/Reels."},
			{Role: "user", Content: "Video: https://www.youtube.com/watch?v=SXN0Y4UN-Q8"},
		},
	})
	if err != nil {
		log.Printf("Error calling OpenAI API: %v\n", err)
	}
	text := resp.Choices[0].Message.Content
	return strings.Split(text, "\n"), nil
}

// Splits video using FFmpeg
func splitVideo(filePath string, keyMoments []string) ([]string, error) {
	var clipPaths []string

	for i, moment := range keyMoments {
		start := strings.TrimSpace(moment)
		outputPath := fmt.Sprintf("uploads/clip_%d.mp4", i+1)

		cmd := exec.Command("ffmpeg", "-i", filePath, "-ss", start, "-t", "30", "-c:v", "libx264", "-preset", "ultrafast", outputPath)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Println("FFmpeg Error:", out.String())
			return nil, err
		}

		clipPaths = append(clipPaths, outputPath)
	}

	return clipPaths, nil
}
func addSubtitlesToVideo(videoFilePath string) (string, error) {
	// Write SRT content to a temporary file
	var srtContent string
	srtContent += fmt.Sprintf("%d\n%s --> %s\n%s\n\n", 5, fmt.Sprintf("00:00:%02d,000", 1), fmt.Sprintf("00:00:%02d,000", 12),"TEST TEXT")
	srtContent += fmt.Sprintf("%d\n%s --> %s\n%s\n\n", 5, fmt.Sprintf("00:00:%02d,000", 12), fmt.Sprintf("00:00:%02d,000", 17),"NEXT TEST TEXT")
	srtFilePath := "subtitles.srt"
	
	err := os.WriteFile(srtFilePath, []byte(srtContent), 0644)
	if err != nil {
		return "", err
	}

	

	videoFilePath_ := "uploads/" + videoFilePath;
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	encryptedName := "output_with_subtitles" + hex.EncodeToString(randomBytes)+ ".mp4"
	outputFilePath := "uploads/" + encryptedName
	fontDir := "/fonts/default"           // Directory containing your font files
	fontName := "sixtyfour"    // Font name as recognized by FFmpeg
	fontSize := 18               // Desired font size
	fontColor := "D20A2E"
	forceStyle := fmt.Sprintf("FontName=%s,FontSize=%d,PrimaryColour=&H%s&", fontName, fontSize, fontColor)

	// Build the FFmpeg command with subtitle styling options
	cmd := exec.Command("ffmpeg", "-i", videoFilePath_,
		"-vf", fmt.Sprintf("subtitles=%s:fontsdir=%s:force_style='%s'", srtFilePath, fontDir, forceStyle),
		"-c:a", "copy", outputFilePath)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		log.Println("Error adding subtitles:", out.String())
		return "", err
	}


	go func(filename string) {
		time.Sleep(10 * time.Minute)
		fileStorage.Lock()
		delete(fileStorage.files, filename)
		fileStorage.Unlock()
		os.Remove("uploads/" + filename)
	}(encryptedName)
	return outputFilePath, nil
}
