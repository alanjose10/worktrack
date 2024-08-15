package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	// command.AddCommand(buildListCommand(app))
	// command.AddCommand(buildUpdateCommand(app))
	// command.AddCommand(buildDeleteCommand(app))

	// command.AddCommand(buildAddNoteCommand(app))

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
			return app.workModel.Insert(args[0], date)
		},
	}

	command.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "Add work for yesterday")
	command.Flags().StringVarP(&dateStr, "date", "d", "", "Add work for a specific date (dd-mm-yyyy)")

	command.MarkFlagsMutuallyExclusive("yesterday", "date")

	return command
}

// func buildListCommand(app *application) *cobra.Command {

// 	command := &cobra.Command{
// 		Use:   "list",
// 		Short: "List all tasks",
// 		Args:  cobra.NoArgs,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			tasks, err := app.taskModel.List()
// 			if err != nil {
// 				return err
// 			}

// 			if len(tasks) == 0 {
// 				fmt.Println(ui.TextError("No tasks found! Add a task using `worktrack add` and try again."))
// 				return nil
// 			}

// 			columns := []string{"ID", "Name", "Project", "Status", "Notes", "Created At", "Updated At"}
// 			var rows [][]string

// 			for _, task := range tasks {
// 				notes, err := app.notesModel.ListByTaskId(task.ID)
// 				if err != nil {
// 					notes = []models.Note{}
// 				}
// 				rows = append(rows, []string{
// 					fmt.Sprintf("%d", task.ID),
// 					task.Name,
// 					task.Project,
// 					task.Status,
// 					fmt.Sprintf("%d", len(notes)),
// 					helpers.GetHumanDate(task.Created),
// 					helpers.GetHumanDate(task.Updated),
// 				})
// 			}

// 			t := ui.TasksTable(columns, rows)
// 			fmt.Println(t)
// 			return nil
// 		},
// 	}

// 	return command
// }

// func buildUpdateCommand(app *application) *cobra.Command {

// 	var project string
// 	var name string
// 	var status int

// 	command := &cobra.Command{
// 		Use:   "update ID",
// 		Short: "Update a task by ID",
// 		Args:  cobra.ExactArgs(1),
// 		PreRunE: func(cmd *cobra.Command, args []string) error {

// 			if status != -1 {
// 				if ok := models.IsValidState(status); !ok {
// 					fmt.Println(ui.TextError(fmt.Sprintf("Invalid status: %v. must be one of (0 - todo, 1 - in progress, 2 - done, 3 - blocked)", status)))
// 					return errors.New("invalid status")
// 				}
// 			}

// 			return nil
// 		},
// 		RunE: func(cmd *cobra.Command, args []string) error {

// 			id, err := strconv.Atoi(args[0])
// 			if err != nil {
// 				return err
// 			}

// 			task, err := app.taskModel.GetById(id)
// 			if err != nil {
// 				if err == models.ErrTaskNotFound {
// 					fmt.Println(ui.TextError(fmt.Sprintf("Could not find task #%d", id)))
// 				}
// 				return err

// 			}

// 			if project != "" {
// 				task.Project = project
// 			}

// 			if name != "" {
// 				task.Name = name
// 			}

// 			if status != -1 {
// 				task.Status = models.State(status).String()
// 			}

// 			return app.taskModel.Update(task)
// 		},
// 	}

// 	command.Flags().StringVarP(&project, "project", "p", "", "Name of the project")
// 	command.Flags().StringVarP(&name, "name", "n", "", "Name of the task")
// 	command.Flags().IntVarP(&status, "status", "s", -1, "Status of the task [0 - todo, 1 - in progress, 2 - done, 3 - blocked]")

// 	return command
// }

// func buildDeleteCommand(app *application) *cobra.Command {

// 	command := &cobra.Command{
// 		Use:   "delete ID",
// 		Short: "Delete a task by ID",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {

// 			id, err := strconv.Atoi(args[0])
// 			if err != nil {
// 				return err
// 			}

// 			return app.taskModel.Delete(id)
// 		},
// 	}

// 	return command
// }

// func buildAddNoteCommand(app *application) *cobra.Command {

// 	var taskId int

// 	command := &cobra.Command{
// 		Use:   "note ID CONTENT",
// 		Short: "Add a note to a task",
// 		Args:  cobra.ExactArgs(2),
// 		PreRunE: func(cmd *cobra.Command, args []string) error {
// 			if id, err := strconv.Atoi(args[0]); err != nil {
// 				return err
// 			} else {
// 				taskId = id
// 			}
// 			return nil
// 		},
// 		RunE: func(cmd *cobra.Command, args []string) error {

// 			task, err := app.taskModel.GetById(taskId)
// 			if err != nil {
// 				if err == models.ErrTaskNotFound {
// 					fmt.Println(ui.TextError(fmt.Sprintf("Could not find task #%d", taskId)))
// 				}
// 				return err
// 			}

// 			err = app.notesModel.Insert(task.ID, args[1])
// 			if err != nil {
// 				return err
// 			}

// 			return nil
// 		},
// 	}
// 	return command
// }

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

			buff := bufio.NewWriter(os.Stdout)
			defer buff.Flush()

			today := helpers.GetCurrentDate()

			fmt.Fprintf(buff, "Standup Report | %s\n\n", today.Format("Jan 2 2006 Monday"))

			days := helpers.GetNPrevWorkingDays(today, goBack)

			for _, d := range days {

				if d.Day() == today.Day() {
					fmt.Fprintf(buff, "%s (Today):\n", d.Format("Jan 2 2006 Monday"))
				} else if d.Day() == today.Day()-1 {
					fmt.Fprintf(buff, "%s (Yesterday):\n", d.Format("Jan 2 2006 Monday"))
				} else {
					fmt.Fprintf(buff, "%s:\n", d.Format("Jan 2 2006 Monday"))
				}

				fmt.Fprintf(buff, "-------------------------------------------\n\n")

				works, err := app.workModel.List(d)
				if err != nil {
					return err
				}

				if len(works) == 0 {
					continue
				} else {
					for _, work := range works {
						fmt.Fprintf(buff, "-> %s\n", work.Content)
					}
				}

				fmt.Fprintf(buff, "\n\n")

			}

			// from := to.AddDate(0, 0, -goBack)

			// for d := from; d.Before(to.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
			// 	fmt.Fprintf(buff, "%s\n", d.Format("Jan 2 2006 Monday"))
			// }

			// fmt.Fprintf(buff, "%s to %s\n\n\n", from.Format("Jan 2 2006 Monday"), to.Format("Jan 2 2006 Monday"))

			// tasks, err := app.taskModel.ListByDateRange(from, to)
			// if err != nil {
			// 	return err
			// }
			// for _, task := range tasks {

			// 	notes, err := app.notesModel.ListByTaskId(task.ID)
			// 	if err != nil {
			// 		notes = []models.Note{}
			// 	}

			// 	fmt.Fprintf(buff, "Task #%d: %s\n", task.ID, task.Name)
			// 	fmt.Fprintf(buff, "Project: %s\n", task.Project)
			// 	fmt.Fprintf(buff, "Status: %s\n", task.Status)
			// 	fmt.Fprintf(buff, "Created: %s\n", helpers.GetHumanDate(task.Created))
			// 	fmt.Fprintf(buff, "Updated: %s\n", helpers.GetHumanDate(task.Updated))
			// 	fmt.Fprintf(buff, "Notes: %d\n", len(notes))

			// 	fmt.Fprintf(buff, "\n")
			// }

			return nil
		},
	}

	command.Flags().IntVarP(&goBack, "back", "b", 2, "Number of days to go back")

	return command
}
