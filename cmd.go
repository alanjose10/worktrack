package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/models"
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
	command.AddCommand(buildListCommand(app))
	command.AddCommand(buildUpdateCommand(app))
	command.AddCommand(buildDeleteCommand(app))

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

	var project string

	command := &cobra.Command{
		Use:   "add NAME",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.taskModel.Insert(args[0], project, 0)
		},
	}

	command.Flags().StringVarP(&project, "project", "p", "default", "Name of the project")

	return command
}

func buildListCommand(app *application) *cobra.Command {

	command := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := app.taskModel.List()
			if err != nil {
				return err
			}

			if len(tasks) == 0 {
				fmt.Println(ui.TextError("No tasks found! Add a task using `worktrack add` and try again."))
				return nil
			}

			columns := []string{"ID", "Name", "Project", "Status", "Created At", "Updated At"}
			var rows [][]string

			for _, task := range tasks {
				rows = append(rows, []string{
					fmt.Sprintf("%d", task.ID),
					task.Name,
					task.Project,
					task.Status,
					helpers.GetHumanDate(task.Created),
					helpers.GetHumanDate(task.Updated),
				})
			}

			t := ui.TasksTable(columns, rows)
			fmt.Println(t)
			return nil
		},
	}

	return command
}

func buildUpdateCommand(app *application) *cobra.Command {

	var project string
	var name string
	var status int

	command := &cobra.Command{
		Use:   "update ID",
		Short: "Update a task by ID",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {

			if status != -1 {
				if ok := models.IsValidState(status); !ok {
					fmt.Println(ui.TextError(fmt.Sprintf("Invalid status: %v. must be one of (0 - todo, 1 - in progress, 2 - done, 3 - blocked)", status)))
					return errors.New("invalid status")
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			task, err := app.taskModel.GetById(id)
			if err != nil {
				if err == models.ErrTaskNotFound {
					fmt.Println(ui.TextError(fmt.Sprintf("Could not find task #%d", id)))
				}
				return err

			}

			if project != "" {
				task.Project = project
			}

			if name != "" {
				task.Name = name
			}

			if status != -1 {
				task.Status = models.State(status).String()
			}

			return app.taskModel.Update(task)
		},
	}

	command.Flags().StringVarP(&project, "project", "p", "", "Name of the project")
	command.Flags().StringVarP(&name, "name", "n", "", "Name of the task")
	command.Flags().IntVarP(&status, "status", "s", -1, "Status of the task [0 - todo, 1 - in progress, 2 - done, 3 - blocked]")

	return command
}

func buildDeleteCommand(app *application) *cobra.Command {

	command := &cobra.Command{
		Use:   "delete ID",
		Short: "Delete a task by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			return app.taskModel.Delete(id)
		},
	}

	return command
}
