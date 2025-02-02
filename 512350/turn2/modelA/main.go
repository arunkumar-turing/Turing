package main

import (
	"fmt"
	"os"

	"github.com/shahrukhq/metadata-extractor/metadataExtractor"
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
	printMetadata(filePath, metadata)
}

func extractMetadata(filePath string) (*metadataExtractor.FileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return metadataExtractor.GetFileInfo(file)
}

func printMetadata(filePath string, metadata *metadataExtractor.FileInfo) {
	fmt.Printf("File Path: %s\n", filePath)
	fmt.Printf("File Size: %d bytes\n", metadata.Size)
	fmt.Printf("File Type: %s\n", metadata.Type)
	fmt.Printf("Modification Time: %s\n", metadata.ModTime)
	fmt.Printf("Permissions: %s\n", metadata.Mode)
}
