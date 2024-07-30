package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/spf13/viper"
)

func (app *Application) GetConfigFileValueByKey(key string) {
	value := viper.GetString(key)
	if value == "" {
		fmt.Printf("%s not found in config file\n", key)
	} else {
		fmt.Printf("%s: %s\n", key, viper.GetString(key))
	}
}

func (app *Application) InitialiseConfigFile() {

	configFilePath := helpers.GetConfigFilePath()

	// Create config file if it does not exist
	helpers.CreateFileIfNotExists(configFilePath)

	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Log level? [%s] (Allowed values are debug, info, warning, error, fatal): ", viper.GetString("log.level"))
		logLevel, _ := reader.ReadString('\n')
		logLevel = strings.TrimSpace(logLevel)
		if logLevel == "" {
			break
		} else {
			if logLevelIsValid(logLevel) {
				viper.Set("log.level", logLevel)
				break
			} else {
				fmt.Printf("%s is invalid. Please enter a valid value\n", logLevel)
			}
		}
	}

	for {
		fmt.Printf("Start date of any sprint? [%s] (DD-MM-YYYY): ", viper.GetString("sprint.start_date"))
		sprintStartDate, _ := reader.ReadString('\n')
		sprintStartDate = strings.TrimSpace(sprintStartDate)
		if sprintStartDate == "" {
			break
		} else {
			if sprintStartDateIsValid(sprintStartDate) {
				viper.Set("sprint.start_date", sprintStartDate)
				break
			} else {
				fmt.Printf("%s is invalid. Please enter a valid value\n", sprintStartDate)
			}
		}
	}

	for {
		fmt.Printf("Number of days in a sprint? [%s] (in days): ", viper.GetString("sprint.duration"))
		sprintDuration, _ := reader.ReadString('\n')
		sprintDuration = strings.TrimSpace(sprintDuration)
		if sprintDuration == "" {
			break
		} else {
			if sprintDurationIsValid(sprintDuration) {
				viper.Set("sprint.duration", sprintDuration)
				break
			} else {
				fmt.Printf("%s is invalid. Please enter a valid value\n", sprintDuration)
			}
		}
	}

	for {
		fmt.Printf("Frequency of standup meetings? [%s] (in days): ", viper.GetString("standup.frequency"))
		standupFrequency, _ := reader.ReadString('\n')
		standupFrequency = strings.TrimSpace(standupFrequency)
		if standupFrequency == "" {
			break
		} else {
			if standupFrequencyIsValid(standupFrequency) {
				viper.Set("standup.frequency", standupFrequency)
				break
			} else {
				fmt.Printf("%s is invalid. Please enter a valid value\n", standupFrequency)
			}
		}
	}

	if err := viper.WriteConfig(); err != nil {
		panic(err)
	} else {
		fmt.Println("Config file created/updated at", configFilePath)
	}

}

func logLevelIsValid(logLevel string) bool {
	switch logLevel {
	case "debug", "info", "warning", "error", "fatal":
		return true
	}
	return false
}

func sprintStartDateIsValid(sprintStartDate string) bool {
	d, err := time.Parse("02-01-2006", sprintStartDate)
	if err != nil {
		return false
	}
	if d.Unix() > time.Now().Unix() {
		return false
	}
	return true
}

func sprintDurationIsValid(sprintDuration string) bool {
	if n, ok := helpers.IsNumber(sprintDuration); ok {
		if helpers.NumberIsInRange(n, 5, 20) {
			return true
		}
	}
	return false
}

func standupFrequencyIsValid(standupFrequency string) bool {
	if n, ok := helpers.IsNumber(standupFrequency); ok {
		if helpers.NumberIsInRange(n, 1, 7) {
			return true
		}
	}
	return false
}
