package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func main() {
	lang := flag.String("lang", "Russian", "Language for Whisper (default Russian)")
	dir := flag.String("dir", ".", "Directory containing audio/video files (default current directory)")
	outputDir := flag.String("output_dir", ".", "Directory to save .srt files (default current directory)")
	model := flag.String("model", "medium", "Whisper model to use (default medium)")
	flag.Parse()

	allowedExt := []string{
		".3gp", ".aac", ".aiff", ".caf", ".flac", ".m4a", ".mka",
		".mp3", ".mp4", ".mpeg", ".mpg", ".mov", ".ogg", ".opus", ".wav", ".webm",
	}

	files, err := os.ReadDir(*dir)
	if err != nil {
		log.Fatalf("Failed to read directory %s: %v", *dir, err)
	}

	var audioFiles []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if !contains(allowedExt, ext) {
			continue
		}
		srtFile := filepath.Join(*outputDir, strings.TrimSuffix(file.Name(), ext)+".srt")
		if _, err := os.Stat(srtFile); err == nil {
			continue
		}
		audioFiles = append(audioFiles, file.Name())
	}

	total := len(audioFiles)
	if total == 0 {
		fmt.Println("No new files to process.")
		return
	}

	for i, filename := range audioFiles {
		filePath := filepath.Join(*dir, filename)
		ext := strings.ToLower(filepath.Ext(filename))
		srtFile := filepath.Join(*outputDir, strings.TrimSuffix(filename, ext)+".srt")

		fmt.Printf("Processing file (%d/%d): %s\n", i+1, total, filename)

		cmd := exec.Command("whisper", filePath,
			"--language", *lang,
			"--model", *model,
			"--output_dir", *outputDir,
			"--output_format", "srt",
		)

		cmd.Stdout = nil
		cmd.Stderr = nil

		if err := cmd.Run(); err != nil {
			log.Printf("Error processing %s: %v", filename, err)
			continue
		}

		fmt.Printf("%s\n", color.New(color.FgHiGreen).Sprintf("Done: %s", srtFile))
	}
}

func contains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}
