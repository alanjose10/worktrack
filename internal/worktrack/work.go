package worktrack

import "fmt"

func (app *App) AddWork(group, content string) error {
	fmt.Fprintf(app.Out, "%s: %s\n", group, content)
	return nil
}
