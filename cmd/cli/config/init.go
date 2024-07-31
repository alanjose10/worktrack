package config

import (
	"fmt"

	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func buildInitCommand(app *worktrack.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "init",
		Short: "Initialise worktrack configuration and save to file at $HOME/.worktrack/config.yaml",
		Long:  `Initialise worktrack configuration and save to file at $HOME/.worktrack/config.yaml.`,
		Run: func(cmd *cobra.Command, args []string) {
			app.InitialiseConfig()
			fmt.Fprintln(app.Out)
			app.GetConfigFile()
		},
	}
	return command
}
