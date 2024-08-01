package items

import (
	"encoding/json"
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

func (work *WorkItem) Add(filePath string) error {

	file, err := file.Get(helpers.GetTimeFromUnix(work.Timestamp), "work.json")
	if err != nil {
		return err
	}
	data, err := file.Read()
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
	if err := file.Write(data); err != nil {
		return err
	}
	return nil
}

// Write function to read all work files from a given date

// Write function to read all work files from a given date range
