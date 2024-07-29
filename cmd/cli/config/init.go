package config

import (
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise worktrack configuration and save to file at $HOME/.worktrack/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {

		logger.Info("Initializing worktrack configuration")
		// Get home directory of user
		// home, err := os.UserHomeDir()
		// if err != nil {
		// 	panic(err)
		// }
		// logger.Debug("Home directory: ", home)

		// configDir := filepath.Join(home, ".worktrack")
		// configFile := filepath.Join(configDir, "config.yaml")

		// if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		// 	panic(err)
		// }

		// _, err = os.Stat(configFile)
		// if err != nil {
		// 	logger.Debug("Creating config file at: ", configFile)
		// 	file, err := os.Create(configFile)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	defer file.Close()
		// } else {
		// 	viper.SetConfigFile(configFile)
		// 	err = viper.ReadInConfig()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// }

		// reader := bufio.NewReader(os.Stdin)

		// fmt.Printf("Log level (debug, info, warning, error, fatal)(%s): ", viper.GetString("log.level"))
		// logLevel, _ := reader.ReadString('\n')

		// fmt.Print("Enter log path: ")
		// logPath, _ := reader.ReadString('\n')

		// fmt.Print("Enter sprint start date (YYYY-MM-DD): ")
		// sprintStartDate, _ := reader.ReadString('\n')

		// fmt.Print("Enter sprint duration (days): ")
		// sprintDuration, _ := reader.ReadString('\n')

		// fmt.Print("Enter standup frequency (days): ")
		// standupFrequency, _ := reader.ReadString('\n')

		// viper.Set("log.level", strings.TrimSpace(logLevel))
		// viper.Set("log.path", strings.TrimSpace(logPath))
		// viper.Set("sprint.start_date", strings.TrimSpace(sprintStartDate))
		// viper.Set("sprint.duration", strings.TrimSpace(sprintDuration))
		// viper.Set("standup.frequency", strings.TrimSpace(standupFrequency))

		// if err := viper.WriteConfig(); err != nil {
		// 	fmt.Println("Error writing config file:", err)
		// 	os.Exit(1)
		// }
		// fmt.Println("Config file created at", configFile)
	},
}

func init() {
	ConfigCmd.AddCommand(initCmd)
}
