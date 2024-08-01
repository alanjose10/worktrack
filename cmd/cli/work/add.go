package work

import (
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func BuildAddCommand(app *worktrack.App) *cobra.Command {
	command := &cobra.Command{
		Use:       "add",
		Short:     "Add a new work entry",
		Long:      `Add a new work entry`,
		ValidArgs: []string{"group", "content"},
		Args:      cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				group   string
				content string
			)

			if len(args) == 2 {
				group = args[0]
				content = args[1]
			} else {
				group = "default"
				content = args[0]
			}

			if err := app.AddWork(group, content); err != nil {
				logger.Fatal(err)
			}
		},
	}

	return command
}
