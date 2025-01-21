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
	IsHidden    bool
}

func GetFileInfo(file *os.File) (*FileInfo, error) {
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileMode := stat.Mode()
	isHidden := fileMode.Name()[0] == '-'

	return &FileInfo{
		Size:        stat.Size(),
		Type:        getFileType(stat),
		ModTime:     stat.ModTime(),
		Permissions: fileMode,
		IsHidden:    isHidden,
	}, nil
}

func getFileType(stat os.FileInfo) string {
	if stat.IsDir() {
		return "directory"
	}
	return "file"
}
