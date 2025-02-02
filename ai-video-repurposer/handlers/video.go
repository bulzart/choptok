package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type VideoStream struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type FFProbeOutput struct {
	Streams []VideoStream `json:"streams"`
}

func getVideoDimensions(videoPath string) (int, error) {
	// Construct the ffprobe command
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-select_streams", "v:0",
		"-show_entries", "stream=width,height",
		"-of", "json",
		videoPath,
	)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, fmt.Errorf("failed to execute ffprobe: %w", err)
	}

	// Parse the JSON output
	var ffprobeOutput FFProbeOutput
	err = json.Unmarshal(out.Bytes(), &ffprobeOutput)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ffprobe output: %w", err)
	}

	// Ensure there is at least one stream
	if len(ffprobeOutput.Streams) == 0 {
		return 0, fmt.Errorf("no video streams found in file")
	}

	return ffprobeOutput.Streams[0].Height, nil
}

func main() {
	videoPath := "path/to/your/video.mp4"
	height, err := getVideoDimensions(videoPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Video dimensions: %dx", height)
}
