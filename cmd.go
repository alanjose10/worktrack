package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
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
	command.AddCommand(buildVersionCommand(app))
	command.AddCommand(buildAddCommand(app))
	command.AddCommand(buildListCommand(app))
	command.AddCommand(buildCleanupCommand(app))

	command.AddCommand(buildTodoCommand(app))

	// command.AddCommand(buildStandupCommand(app))

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

// Need a command to cleanup the database of all existing data
func buildCleanupCommand(app *application) *cobra.Command {
	var cleanWork, cleanTodo, cleanBlocker bool

	command := &cobra.Command{
		Use:   "cleanup",
		Short: "Clean up the database of all existing data",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !cleanWork && !cleanTodo && !cleanBlocker {
				cleanWork, cleanTodo, cleanBlocker = true, true, true
			}

			var confirm bool
			prompt := huh.NewConfirm().
				Description("This action cannot be undone.").
				Value(&confirm).
				TitleFunc(func() string {
					title := "Are you sure you want to delete the selected data?"
					title += " ("
					if cleanWork {
						title += "WORK, "
					}
					if cleanTodo {
						title += "TODOS, "
					}
					if cleanBlocker {
						title += "BLOCKERS, "
					}
					title = strings.TrimSuffix(title, ", ")
					title += ")"
					return title
				}, nil).
				Affirmative("Yes, delete").
				Negative("No, cancel")

			err := prompt.Run()
			if err != nil {
				return err
			}

			if !confirm {
				fmt.Println("Cleanup cancelled.")
				return nil
			}

			if cleanWork {
				err = app.workModel.DeleteAll()
				if err != nil {
					return fmt.Errorf("failed to delete work items: %w", err)
				}
				fmt.Println("Work items deleted successfully.")
			}

			if cleanTodo {
				err = app.todoModel.DeleteAll()
				if err != nil {
					return fmt.Errorf("failed to delete todo items: %w", err)
				}
				fmt.Println("Todo items deleted successfully.")
			}

			if cleanBlocker {
				err = app.blockerModel.DeleteAll()
				if err != nil {
					return fmt.Errorf("failed to delete blocker items: %w", err)
				}
				fmt.Println("Blocker items deleted successfully.")
			}

			return nil
		},
	}

	command.Flags().BoolVarP(&cleanWork, "work", "w", false, "Clean up work items")
	command.Flags().BoolVarP(&cleanTodo, "todo", "t", false, "Clean up todo items")
	command.Flags().BoolVarP(&cleanBlocker, "blocker", "b", false, "Clean up blocker items")

	command.MarkFlagsOneRequired("work", "todo", "blocker")

	return command
}

func buildVersionCommand(app *application) *cobra.Command {

	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version of the application",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Println(app.version)
			return err
		},
	}

	return command
}

// buildAddCommand creates a new cobra command to add a work, todo or blocker item
func buildAddCommand(app *application) *cobra.Command {

	var date time.Time

	var yesterday bool
	var dateStr string

	// var todo bool
	// var blocker bool

	var item = "work"
	var promptInput = false

	command := &cobra.Command{
		Use:   "add [todo|blocker] [text]",
		Short: "Add items",
		Args:  cobra.RangeArgs(0, 2),
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

			if len(args) == 0 {
				item = "work"
				promptInput = true
			} else if len(args) == 1 {
				switch args[0] {
				case "todo", "todos", "t":
					{
						item = "todo"
						promptInput = true
					}
				case "blocker", "blockers", "block", "b":
					{
						item = "blocker"
						promptInput = true
					}
				default:
					{
						item = "work"
						promptInput = false
					}
				}
			} else if len(args) == 2 {

				switch args[0] {
				case "todo", "todos", "t":
					{
						item = "todo"
						promptInput = false
					}
				case "blocker", "blockers", "block", "b":
					{
						item = "blocker"
						promptInput = false
					}
				default:
					return fmt.Errorf("invalid item type %s", args[0])
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			var textInput string
			if promptInput {
				err := huh.NewText().
					TitleFunc(func() string {

						switch item {
						case "todo":
							return "Add a todo item"
						case "blocker":
							return "Add a blocker item"
						}
						return "Add a work item"
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

				if item == "todo" || item == "blocker" {
					textInput = args[1]
				} else {
					if args[0] == "help" {
						return cmd.Help()
					} else {
						textInput = args[0]
					}
				}
			}

			switch item {
			case "todo":
				return app.todoModel.Insert(textInput, date)

			case "blocker":
				return app.blockerModel.Insert(textInput, date)
			}
			return app.workModel.Insert(textInput, date)
		},
		Example: addCmdExamples,
	}

	command.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "Add work for yesterday")
	command.Flags().StringVarP(&dateStr, "date", "d", "", "Add work for a specific date (dd-mm-yyyy)")
	command.MarkFlagsMutuallyExclusive("yesterday", "date")

	return command
}

func buildListCommand(app *application) *cobra.Command {
	var (
		fromDate   time.Time
		d, w, m, y int
	)

	command := &cobra.Command{
		Use:     "list [todo|blocker] -d|-w|-m|-y",
		Short:   "List items",
		Example: listCmdExamples,
		Args:    cobra.RangeArgs(0, 1),
		PreRunE: func(cmd *cobra.Command, args []string) error {

			if d == 0 && w == 0 && m == 0 && y == 0 {
				d = 7
			}

			switch {
			case d > 0:
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(0, 0, -d))
			case w > 0:
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(0, 0, -w*7))
			case m > 0:
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(0, -m, 0))
			case y > 0:
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(-y, 0, 0))
			default:
				return fmt.Errorf("no time range specified")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			toDate := helpers.GetDateCeil(helpers.GetCurrentDate())

			if len(args) == 1 {
				switch args[0] {
				case "todo", "todos", "t":
					return app.listTodo(fromDate, toDate)
				case "blocker", "blockers", "block", "b":
					return app.listBlocker(fromDate, toDate)
				default:
					return fmt.Errorf("invalid item type %s", args[0])
				}
			}
			return app.listWork(fromDate, toDate)
		},
	}

	command.Flags().IntVarP(&d, "days", "d", 0, "Go back n days")
	command.Flags().IntVarP(&w, "weeks", "w", 0, "Go back n weeks")
	command.Flags().IntVarP(&m, "months", "m", 0, "Go back n months")
	command.Flags().IntVarP(&y, "years", "y", 0, "Go back n years")

	command.MarkFlagsMutuallyExclusive("days", "weeks", "months", "years")

	return command
}

func buildTodoCommand(app *application) *cobra.Command {

	var (
		do       bool
		undo     bool
		fromDate time.Time
		d        int
		w        int
		m        int
		y        int
	)

	command := &cobra.Command{
		Use:   "todo --do|--undo -d|-w|-m|-y",
		Short: "Mark a todo item as done or undone",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {

			if d == 0 && w == 0 && m == 0 && y == 0 {
				d = 7
			}

			if d > 0 {
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(0, 0, -1*d))
			} else if w > 0 {
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(0, 0, -1*w*7))
			} else if m > 0 {
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(0, -1*m, 0))
			} else if y > 0 {
				fromDate = helpers.GetDateFloor(helpers.GetCurrentDate().AddDate(-1*y, 0, 0))
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			toDate := helpers.GetDateCeil(helpers.GetCurrentDate())

			if do {
				app.doTodo(fromDate, toDate)
			} else if undo {
				app.undoTodo(fromDate, toDate)
			}
		},
	}

	command.Flags().BoolVar(&do, "do", false, "Mark a pending todo item as done")
	command.Flags().BoolVar(&undo, "undo", false, "Mark a compoeted todo item as undone")

	command.MarkFlagsOneRequired("do", "undo")
	command.MarkFlagsMutuallyExclusive("do", "undo")

	command.Flags().IntVarP(&d, "days", "d", 0, "Go back n days")
	command.Flags().IntVarP(&w, "weeks", "w", 0, "Go back n weeks")
	command.Flags().IntVarP(&m, "months", "m", 0, "Go back n months")
	command.Flags().IntVarP(&y, "years", "y", 0, "Go back n years")

	command.MarkFlagsMutuallyExclusive("days", "weeks", "months", "years")

	return command
}

// func buildStandupCommand(app *application) *cobra.Command {

// 	var goBack int

// 	command := &cobra.Command{
// 		Use:   "standup",
// 		Short: "Get a standup report",
// 		Args:  cobra.NoArgs,
// 		PreRunE: func(cmd *cobra.Command, args []string) error {
// 			if goBack < 1 || goBack > 7 {
// 				fmt.Println(ui.TextError("Invalid number of days to go back. Must be between 0 and 7"))
// 				return errors.New("invalid number of days")
// 			}
// 			return nil
// 		},
// 		RunE: func(cmd *cobra.Command, args []string) error {

// 			fmt.Println(app.printStandupReport(goBack))

// 			return nil
// 		},
// 	}

// 	command.Flags().IntVarP(&goBack, "back", "b", 2, "Number of days to go back")

// 	return command
// }
