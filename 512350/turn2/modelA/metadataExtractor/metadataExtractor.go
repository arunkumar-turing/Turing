package metadataExtractor

import (
    "os"
    "time"
)

type FileInfo struct {
    Size    int64
    Type    string
    ModTime time.Time
    Mode    os.FileMode
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

    return &FileInfo{
        Size:    stat.Size(),
        Type:    fileType,
        ModTime: stat.ModTime(),
        Mode:    stat.Mode(),
    }, nil
}
