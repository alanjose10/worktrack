package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use: "config",
}

func init() {
	ConfigCmd.AddCommand(viewCmd)
	ConfigCmd.AddCommand(initCmd)
}
