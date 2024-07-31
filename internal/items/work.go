package items

import (
	"fmt"
	"time"

	"github.com/alanjose10/worktrack/internal/file"
	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/logger"
)

type WorkItem struct {
	Group     string `json:"group"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

func NewWork(group, content string, ts time.Time) *WorkItem {
	return &WorkItem{
		Group:     group,
		Content:   content,
		Timestamp: ts.Unix(),
	}
}

func (work *WorkItem) JsonItem() map[string]interface{} {
	return map[string]interface{}{
		"group":     work.Group,
		"content":   work.Content,
		"timestamp": work.Timestamp,
	}
}

func (work *WorkItem) Add(filePath string) error {

	file, err := file.Get(helpers.GetTimeFromUnix(work.Timestamp), "work.json")
	if err != nil {
		return err
	}
	logger.Debug(fmt.Sprintf("%+v", file))

	return nil
}

// Write function to read all work files from a given date

// Write function to read all work files from a given date range
