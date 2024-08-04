package work

import (
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func BuildGetCommand(app *worktrack.App) *cobra.Command {
	var ts time.Time
	var date string
	var yesterday bool
	var today bool

	command := &cobra.Command{
		Use:   "get",
		Short: "Get worked items",
		PreRunE: func(cmd *cobra.Command, args []string) error {

			if today {
				ts = helpers.GetCurrentDate()
			} else if yesterday {
				ts = helpers.GetYesterdayDate()
			} else {
				var err error
				ts, err = helpers.ParseDate(date)
				if err != nil {
					return err
				}
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			if err := app.GetWork(ts); err != nil {
				logger.Fatal(err)
			}
		},
	}

	command.Flags().StringVarP(&date, "date", "d", "", "Worked date")

	command.Flags().BoolVarP(&today, "today", "t", false, "Get today's work")

	command.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "Get yesterday's work")

	command.MarkFlagsOneRequired("date", "today", "yesterday")
	command.MarkFlagsMutuallyExclusive("date", "today", "yesterday")

	return command
}
