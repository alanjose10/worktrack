package cli

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/alanjose10/worktrack/cmd/cli/config"
	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool

var RootCmd = &cobra.Command{
	Version: "0.0.1",
	Use:     "worktrack",
	Short:   "worktrack is a CLI tool to track your work",
	Long:    `worktrack is a CLI tool to track your work.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

func init() {

	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))

	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(config.ConfigCmd)

}

func initConfig() {

	viper.SetDefault("log.level", "info")
	viper.SetDefault("sprint.start_date", "29-07-2024")
	viper.SetDefault("sprint.duration", "10")
	viper.SetDefault("standup.frequency", "2")

	worktrackDir := helpers.GetWorktrackDir()
	helpers.CreateDirectoryIfNotExists(worktrackDir)

	configFilePath := helpers.GetConfigFilePath()

	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			logger.Warning(fmt.Sprintf("Config file not found at %s", configFilePath))
		} else {
			panic(err)
		}
	}
}
