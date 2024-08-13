package main

import (
	"fmt"

	"github.com/alanjose10/worktrack/internal/components"
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
				fmt.Println(components.TextError("No tasks found! Add a task using `worktrack add` and try again."))
				return nil
			}

			columns := []string{"ID", "Name", "Project", "Status", "Created At"}
			var rows [][]string

			for _, task := range tasks {
				rows = append(rows, []string{
					fmt.Sprintf("%d", task.ID),
					task.Name,
					task.Project,
					task.Status,
					task.Created.Format("Mon Jan 2 2006 at 3:04 PM"),
				})
			}

			t := components.Table(columns, rows)
			fmt.Println(t)
			return nil
		},
	}

	return command
}
