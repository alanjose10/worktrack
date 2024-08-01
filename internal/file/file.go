package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/logger"
)

type File struct {
	Location string
	Name     string
	Path     string
}

func Get(ts time.Time, name string) (*File, error) {
	location := helpers.GetStorageDir(ts)
	if err := helpers.CreateDirectoryIfNotExists(location); err != nil {
		return nil, err
	}
	filePath := filepath.Join(location, name)
	helpers.CreateFileIfNotExists(filePath)

	return &File{
		Location: location,
		Name:     name,
		Path:     filePath,
	}, nil
}

func (file *File) Read() ([]byte, error) {
	logger.Debug(fmt.Sprintf("Reading from %s", file.Path))
	f, err := os.Open(file.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read file content
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (file *File) Write(data []byte) error {
	logger.Debug(fmt.Sprintf("Writing to %s", file.Path))

	f, err := os.OpenFile(file.Path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Clear file content
	f.Truncate(0)
	f.Seek(0, 0)

	// Write to file
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}
