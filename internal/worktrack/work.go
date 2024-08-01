package worktrack

import (
	"time"

	"github.com/alanjose10/worktrack/internal/items"
)

func (app *App) AddWork(group, content string, ts time.Time) error {

	// ts := helpers.GetCurrentDate()
	work := items.NewWork(group, content, ts)

	if err := work.Add("work.json"); err != nil {
		return err
	}

	return nil
}
