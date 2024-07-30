package config

import (
	"log"

	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func buildViewCommand(app *worktrack.App) *cobra.Command {
	var key string
	command := &cobra.Command{
		Use:   "view",
		Short: "View the app configuration file contents at $HOME/.worktrack/config.yaml",
		Long:  `View the app configuration file contents at $HOME/.worktrack/config.yaml.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Hello from worktrack config view")
		},
	}
	command.Flags().StringVarP(&key, "key", "k", "", "Key to view")
	return command
}
