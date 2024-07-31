package worktrack

import (
	"fmt"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/items"
)

func (app *App) AddWork(group, content string) error {
	fmt.Fprintf(app.Out, "%s: %s\n", group, content)

	ts := helpers.GetCurrentDate()
	work := items.NewWork(group, content, ts)

	if err := work.Add("work.json"); err != nil {
		return err
	}

	return nil
}
