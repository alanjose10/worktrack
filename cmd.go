package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/ui"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func buildRootCommand(app *application) *cobra.Command {

	command := &cobra.Command{
		Use:   "worktrack",
		Short: "A CLI work tracker tool to make sprints & standups easier",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	command.AddCommand(buildWhereCommand(app))
	command.AddCommand(buildAddCommand(app))
	command.AddCommand(buildStandupCommand(app))

	command.AddCommand(buildListCommand(app))

	return command

}

func buildWhereCommand(app *application) *cobra.Command {

	command := &cobra.Command{
		Use:   "where",
		Short: "Show where your tasks are stored",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Println(app.dataPath)
			return err
		},
	}

	return command
}

func buildAddCommand(app *application) *cobra.Command {

	var date time.Time

	var yesterday bool
	var dateStr string

	var todo bool
	var blocker bool

	examples := `
Add a work item using inline input

	worktrack add "Worked on documenting the APIs"

Add a work item via prompt (An input prompt will be provided)

	worktrack add

Add a todo item

	worktrack add --todo

Add a blocker item

	worktrack add --blocker

Add an item for yesterday

	worktrack add -y

Add an item for a specific date

	worktrack add --date 20-10-1994
	`

	command := &cobra.Command{
		Use:   "add WORK",
		Short: "Add details of a work",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {

			if yesterday {
				date = helpers.GetYesterdayDate()
			} else if dateStr != "" {
				var err error
				date, err = helpers.ParseDate(dateStr)
				if err != nil {
					return err
				}
			} else {
				date = helpers.GetCurrentDate()
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			var textInput string
			if len(args) == 0 {

				err := huh.NewText().
					TitleFunc(func() string {
						if todo {
							return "Please input the todo item."
						}
						if blocker {
							return "Please enter the blocker details."
						}
						return "Please enter the work item details."
					}, nil).
					Value(&textInput).
					Validate(func(s string) error {
						if strings.TrimSpace(s) == "" {
							fmt.Println("No input provided.")
							return fmt.Errorf("no input provided")
						}
						return nil
					}).
					Run()
				if err != nil {
					os.Exit(1)
				}

			} else {
				textInput = args[0]
			}

			if todo {
				return app.todoModel.Insert(textInput, date)
			}

			if blocker {
				return app.blockerModel.Insert(textInput, date)
			}

			return app.workModel.Insert(textInput, date)
			// return nil
		},
		Example: examples,
	}

	command.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "Add work for yesterday")
	command.Flags().StringVarP(&dateStr, "date", "d", "", "Add work for a specific date (dd-mm-yyyy)")

	command.Flags().BoolVar(&todo, "todo", false, "Add a todo item")
	command.Flags().BoolVar(&blocker, "blocker", false, "Add a blocker item")

	command.MarkFlagsMutuallyExclusive("yesterday", "date")

	command.MarkFlagsMutuallyExclusive("todo", "blocker")

	return command
}

func buildStandupCommand(app *application) *cobra.Command {

	var goBack int

	command := &cobra.Command{
		Use:   "standup",
		Short: "Get a standup report",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if goBack < 1 || goBack > 7 {
				fmt.Println(ui.TextError("Invalid number of days to go back. Must be between 0 and 7"))
				return errors.New("invalid number of days")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println(app.printStandupReport(goBack))

			return nil
		},
	}

	command.Flags().IntVarP(&goBack, "back", "b", 2, "Number of days to go back")

	return command
}

func buildListCommand(app *application) *cobra.Command {

	var (
		date time.Time
		d    int
		w    int
		m    int
		y    int
	)

	command := &cobra.Command{
		Use:   "list",
		Short: "List items",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if d > 0 {
				date = helpers.GetCurrentDate().AddDate(0, 0, -1*d)
			} else if w > 0 {
				date = helpers.GetCurrentDate().AddDate(0, 0, -1*w*7)
			} else if m > 0 {
				date = helpers.GetCurrentDate().AddDate(0, -1*m, 0)
			} else if y > 0 {
				date = helpers.GetCurrentDate().AddDate(-1*y, 0, 0)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println(helpers.GetHumanDate(date))

			return nil
		},
	}

	command.Flags().IntVarP(&d, "days", "d", 0, "Go back n days")
	command.Flags().IntVarP(&w, "weeks", "w", 0, "Go back n weeks")
	command.Flags().IntVarP(&m, "months", "m", 0, "Go back n months")
	command.Flags().IntVarP(&y, "years", "y", 0, "Go back n years")

	command.MarkFlagsOneRequired("days", "weeks", "months", "years")
	command.MarkFlagsMutuallyExclusive("days", "weeks", "months", "years")

	return command
}
