package work

import (
	"errors"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func BuildAddCommand(app *worktrack.App) *cobra.Command {

	var date string
	var ts time.Time
	var yesterday bool

	command := &cobra.Command{
		Use:       "add",
		Short:     "Add a new work entry",
		Long:      `Add a new work entry`,
		ValidArgs: []string{"group", "content"},
		Args:      cobra.RangeArgs(1, 2),
		PreRunE: func(cmd *cobra.Command, args []string) error {

			// Do not allow both date and yesterday flags to be set
			if date != "" && yesterday {
				return errors.New("only one of date or yesterday flags can be set")
			}

			// Validate the date flag
			var err error
			if date != "" {
				logger.Debug(date)
				ts, err = helpers.ParseDate(date)
				if err != nil {
					return err
				}
			} else {
				ts = helpers.GetCurrentDate()
			}

			// If yesterday flag is set, subtract 24 hours from the current date
			if yesterday {
				ts = helpers.GetCurrentDate()
				ts = ts.AddDate(0, 0, -1)
			}

			return nil
		},
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

			if err := app.AddWork(group, content, ts); err != nil {
				logger.Fatal(err)
			}
		},
	}

	command.Flags().StringVarP(&date, "date", "d", "", "Date of the work entry")

	command.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "Add work entry for yesterday")

	return command
}
