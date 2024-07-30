package config

import (
	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func buildViewCommand(app *worktrack.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "view",
		Short: "View the app configuration file contents at $HOME/.worktrack/config.yaml",
		Long:  `View the app configuration file contents at $HOME/.worktrack/config.yaml.`,
		Run: func(cmd *cobra.Command, args []string) {
			app.GetConfigFile()
		},
	}
	return command
}
