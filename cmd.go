package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/ui"
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

	command := &cobra.Command{
		Use:   "add WORK",
		Short: "Add details of a work",
		Args:  cobra.ExactArgs(1),
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

			if todo {
				return app.todoModel.Insert(args[0], date)
			}

			if blocker {
				return app.blockerModel.Insert(args[0], date)
			}

			return app.workModel.Insert(args[0], date)
		},
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

			todos, err := app.todoModel.List()
			if err != nil {
				return err
			}

			blockers, err := app.blockerModel.List()
			if err != nil {
				return err
			}

			fmt.Println(ui.PrintStandupReport(goBack, todos, blockers))

			return nil
		},
	}

	command.Flags().IntVarP(&goBack, "back", "b", 2, "Number of days to go back")

	return command
}
