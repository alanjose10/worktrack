package cli

import (
	"os"
	"path/filepath"

	"github.com/alanjose10/worktrack/cmd/cli/config"
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Version: "0.0.1",
	Use:     "worktrack",
	Short:   "worktrack is a CLI tool to track your work",
	Long: `worktrack is a CLI tool to track your work. 
	
It can help you 
	- Track work done
	- Keep list of todo items
	- Keep list of blockers
	- Generate report for standup
	- Generate report for sprint retrospective
	- Generate yearly report
	
	
Examples:

Initialize worktrack and setup configuration file
worktrack config init

View the configuration
worktrack config view

Add a new work item
worktrack add -t "Website" "Setup authentication using JWT"

`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(config.ConfigCmd)

}

func initConfig() {

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(home, ".worktrack")
	configFile := filepath.Join(configDir, "config.yaml")

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err == nil {
		logger.Debug("Using config file:", viper.ConfigFileUsed())
	}

}
