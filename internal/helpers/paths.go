package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}

func GetWorktrackDir() string {
	return filepath.Join(getHomeDir(), ".worktrack")
}

func GetConfigFilePath() string {
	return filepath.Join(GetWorktrackDir(), "config")
}

func GetStorageDir(ts time.Time) string {
	return filepath.Join(
		GetWorktrackDir(),
		fmt.Sprintf("%d/%d/%d", ts.Year(), ts.Month(), ts.Day()),
	)
}

func CreateDirectoryIfNotExists(dirPath string) error {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func CreateFileIfNotExists(filePath string) {
	_, err := os.Stat(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}
