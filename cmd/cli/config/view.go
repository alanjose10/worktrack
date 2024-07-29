package config

import (
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the app configuration file contents at $HOME/.worktrack/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("View config command called")
		logger.Debug(viper.GetString("sprint.start_date"))
		logger.Debug(viper.GetString("sprint.duration"))

	},
}

func init() {
	ConfigCmd.AddCommand(viewCmd)
}
