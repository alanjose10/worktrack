package file

import (
	"fmt"
	"io"
	"os"

	"github.com/alanjose10/worktrack/internal/logger"
)

func ReadFile(path string) ([]byte, error) {
	logger.Debug(fmt.Sprintf("Reading from %s", path))
	f, err := os.Open(path)
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

func WriteFile(path string, data []byte) error {
	logger.Debug(fmt.Sprintf("Writing to %s", path))

	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
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
