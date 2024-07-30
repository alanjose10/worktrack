package config

import (
	"github.com/alanjose10/worktrack/internal/app"
	"github.com/spf13/cobra"
)

var key string

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the app configuration file contents at $HOME/.worktrack/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		app := &app.Application{}
		app.GetConfigFileValueByKey(key)
	},
}

func init() {
	viewCmd.Flags().StringVarP(&key, "key", "k", "", "Key to view")
}
