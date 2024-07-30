package config

import (
	"github.com/alanjose10/worktrack/internal/app"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise worktrack configuration and save to file at $HOME/.worktrack/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		app := &app.Application{}
		app.InitialiseConfigFile()
	},
}
