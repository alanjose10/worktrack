package config

import (
	"path/filepath"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/logger"
	"github.com/spf13/viper"
)

type Sprint struct {
	StartDate string `mapstructure:"start_date"`
	Duration  int    `mapstructure:"duration"`
}

type Standup struct {
	Frequency int
}

type Config struct {
	LogLevel string  `mapstructure:"log_level"`
	Sprint   Sprint  `mapstructure:"sprint"`
	Standup  Standup `mapstructure:"standup"`
}

func NewConfig() *Config {
	return &Config{
		LogLevel: "info",
		Sprint: Sprint{
			StartDate: "29-07-2024",
			Duration:  10,
		},
		Standup: Standup{
			Frequency: 2,
		},
	}
}

func (c *Config) Load() error {
	dir := helpers.GetWorktrackDir()
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(dir)
	if err := v.ReadInConfig(); err != nil {
		logger.Warning(err.Error())

	}

	if err := v.Unmarshal(c); err != nil {
		logger.Fatal(err)
	}
	return nil
}

func (c *Config) Save() error {

	dir := helpers.GetWorktrackDir()
	helpers.CreateDirectoryIfNotExists(dir)
	helpers.CreateFileIfNotExists(filepath.Join(dir, "config"))

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(dir)

	v.Set("log_level", c.LogLevel)
	v.Set("sprint.start_date", c.Sprint.StartDate)
	v.Set("sprint.duration", c.Sprint.Duration)
	v.Set("standup.frequency", c.Standup.Frequency)

	if err := v.WriteConfig(); err != nil {
		logger.Warning(err.Error())
		return err
	}
	return nil
}
