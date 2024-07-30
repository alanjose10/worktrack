package worktrack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/logger"
)

func (app *App) GetConfigFile() {
	f, err := os.Open(app.Config.Location)
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()
	io.Copy(app.Out, f)
}

func (app *App) InitialiseConfig() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprintf(app.Out, "Log level? [%s] (Allowed values are debug, info, warning, error, fatal): ", app.Config.LogLevel)
		logLevel, _ := reader.ReadString('\n')
		logLevel = strings.TrimSpace(logLevel)
		if logLevel == "" {
			break
		} else {
			if logLevelIsValid(logLevel) {
				app.Config.LogLevel = logLevel
				break
			} else {
				fmt.Fprintf(app.Out, "%s is invalid. Please enter a valid value\n", logLevel)
			}
		}
	}

	for {
		fmt.Fprintf(app.Out, "Start date of any sprint? [%s] (DD-MM-YYYY): ", app.Config.Sprint.StartDate)
		sprintStartDate, _ := reader.ReadString('\n')
		sprintStartDate = strings.TrimSpace(sprintStartDate)
		if sprintStartDate == "" {
			break
		} else {
			if sprintStartDateIsValid(sprintStartDate) {
				app.Config.Sprint.StartDate = sprintStartDate
				break
			} else {
				fmt.Fprintf(app.Out, "%s is invalid. Please enter a valid value\n", sprintStartDate)
			}
		}
	}

	for {
		fmt.Fprintf(app.Out, "Number of days in a sprint? [%d] (in days): ", app.Config.Sprint.Duration)
		sprintDuration, _ := reader.ReadString('\n')
		sprintDuration = strings.TrimSpace(sprintDuration)
		if sprintDuration == "" {
			break
		} else {
			if n, ok := sprintDurationIsValid(sprintDuration); ok {
				app.Config.Sprint.Duration = n
				break
			} else {
				fmt.Fprintf(app.Out, "%s is invalid. Please enter a valid value\n", sprintDuration)
			}
		}
	}

	for {
		fmt.Fprintf(app.Out, "Frequency of standup meetings? [%d] (in days): ", app.Config.Standup.Frequency)
		standupFrequency, _ := reader.ReadString('\n')
		standupFrequency = strings.TrimSpace(standupFrequency)
		if standupFrequency == "" {
			break
		} else {
			if n, ok := standupFrequencyIsValid(standupFrequency); ok {
				app.Config.Standup.Frequency = n
				break
			} else {
				fmt.Fprintf(app.Out, "%s is invalid. Please enter a valid value\n", standupFrequency)
			}
		}
	}

	if err := app.Config.Save(); err != nil {
		logger.Fatal(err)
	} else {
		fmt.Fprintf(app.Out, "Config file at [%s] updated successfully\n", app.Config.Location)
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

func sprintDurationIsValid(sprintDuration string) (int, bool) {
	if n, ok := helpers.IsNumber(sprintDuration); ok {
		if helpers.NumberIsInRange(n, 5, 20) {
			return n, true
		}
	}
	return 0, false
}

func standupFrequencyIsValid(standupFrequency string) (int, bool) {
	if n, ok := helpers.IsNumber(standupFrequency); ok {
		if helpers.NumberIsInRange(n, 1, 7) {
			return n, true
		}
	}
	return 0, false
}
