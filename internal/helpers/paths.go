package helpers

import (
	"os"
	"path/filepath"
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

func CreateDirectoryIfNotExists(dirPath string) {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		panic(err)
	}
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
