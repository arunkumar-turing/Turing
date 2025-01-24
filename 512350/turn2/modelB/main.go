package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	metadata, err := extractMetadata(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("File Path: %s\n", filePath)
	fmt.Printf("File Size: %d bytes\n", metadata.Size)
	fmt.Printf("File Type: %s\n", metadata.Type)
	fmt.Printf("Modification Time: %s\n", metadata.ModTime)
	fmt.Printf("Permissions: %s\n", metadata.Permissions)
}

type FileInfo struct {
	Size        int64
	Type        string
	ModTime     time.Time
	Permissions os.FileMode
}

func extractMetadata(filePath string) (*FileInfo, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	fileType := "file"
	if info.IsDir() {
		fileType = "directory"
	}

	return &FileInfo{
		Size:        info.Size(),
		Type:        fileType,
		ModTime:     info.ModTime(),
		Permissions: info.Mode(),
	}, nil
}
