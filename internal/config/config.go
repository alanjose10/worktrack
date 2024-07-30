package config

type Sprint struct {
	StartDate string
	Duration  int
}

type Standup struct {
	Frequency int
}

type Config struct {
	LogLevel string
	Sprint   Sprint
	Standup  Standup
}

func Load() (*Config, error) {
	return nil, nil
}
