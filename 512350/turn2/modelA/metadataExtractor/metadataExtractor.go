package metadataExtractor

import (
	"os"
	"time"
)

type FileInfo struct {
	Size        int64
	Type        string
	ModTime     time.Time
	Permissions os.FileMode
	Hidden      bool
}

// GetFileInfo extracts metadata from an opened file.
func GetFileInfo(file *os.File) (*FileInfo, error) {

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	var fileType string
	if stat.IsDir() {
		fileType = "directory"
	} else {
		fileType = "file"
	}

	// Determine if the file is hidden
	isHidden := isFileHidden(file.Name())

	return &FileInfo{
		Size:        stat.Size(),
		Type:        fileType,
		ModTime:     stat.ModTime(),
		Permissions: stat.Mode().Perm(), // Get file permissions
		Hidden:      isHidden,
	}, nil
}

// isFileHidden checks if a file is hidden based on its name.
func isFileHidden(name string) bool {
	return len(name) > 1 && name[0] == '.' // Checks if file name starts with '.'
}