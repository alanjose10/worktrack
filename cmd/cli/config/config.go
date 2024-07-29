package config

import (
	"github.com/alanjose10/worktrack/internal/logger"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Initializing worktrack configuration")
	},
}

func init() {
	ConfigCmd.AddCommand(viewCmd)
	ConfigCmd.AddCommand(initCmd)
}
