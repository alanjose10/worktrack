package worktrack

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/items"
)

func (app *App) AddWork(group, content string, ts time.Time) error {

	// ts := helpers.GetCurrentDate()
	work := items.NewWork(group, content, ts)

	if err := work.Add(); err != nil {
		return err
	}

	return nil
}

func (app *App) GetWork(ts time.Time) error {

	workItems, err := items.GetWork(ts)
	if err != nil {
		return err
	}

	fmt.Fprintf(app.Out, "Date: %s\n\n", helpers.GetHumanDate(ts))

	if len(workItems) == 0 {
		fmt.Fprintln(app.Out, "No entries found")
	} else {
		sort.Slice(workItems, func(i, j int) bool {
			res := strings.Compare(workItems[i].Group, workItems[j].Group)
			return res <= 0
		})
		for index, item := range workItems {
			fmt.Fprintf(app.Out, "%d: %s - %s\n", index, item.Group, item.Content)
		}
	}
	return nil
}
