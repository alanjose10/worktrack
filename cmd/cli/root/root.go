package root

import (
	"log"

	"github.com/alanjose10/worktrack/cmd/cli/config"
	"github.com/alanjose10/worktrack/cmd/cli/work"
	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func BuildRootCommand(app *worktrack.App) *cobra.Command {

	command := &cobra.Command{
		Use:   "worktrack",
		Short: "worktrack is a CLI tool to track your work",
		Long:  `worktrack is a CLI tool to track your work.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Hello from worktrack")
		},
	}

	command.AddCommand(config.BuildConfigCommand(app))
	command.AddCommand(work.BuildAddCommand(app))
	command.AddCommand(work.BuildGetCommand(app))
	return command

}
