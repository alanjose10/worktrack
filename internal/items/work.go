package items

import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/alanjose10/worktrack/internal/file"
	"github.com/alanjose10/worktrack/internal/helpers"
)

type WorkItem struct {
	Id        string `json:"id"`
	Group     string `json:"group"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

func NewWork(group, content string, ts time.Time) *WorkItem {
	return &WorkItem{
		Id:        helpers.GetUUID(),
		Group:     group,
		Content:   content,
		Timestamp: ts.Unix(),
	}
}

func (work *WorkItem) Add() error {
	tsUnix := helpers.GetTimeFromUnix(work.Timestamp)
	location := helpers.GetStorageDir(tsUnix)
	if err := helpers.CreateDirectoryIfNotExists(location); err != nil {
		return err
	}
	path := filepath.Join(location, "work.json")
	helpers.CreateFileIfNotExists(path)

	data, err := file.ReadFile(path)
	if err != nil {
		return err
	}

	if helpers.RemoveWhiteSpaces(string(data)) == "" {
		data = []byte("[]")
	}

	var workItems []WorkItem
	err = json.Unmarshal(data, &workItems)
	if err != nil {
		return err
	}

	workItems = append(workItems, *work)

	data, err = json.Marshal(workItems)
	if err != nil {
		return err

	}
	if err := file.WriteFile(path, data); err != nil {
		return err
	}
	return nil
}

// Write function to read all work files from a given date

// Write function to read all work files from a given date range
