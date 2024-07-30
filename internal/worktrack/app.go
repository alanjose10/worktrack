package worktrack

import (
	"io"
	"os"

	"github.com/alanjose10/worktrack/internal/config"
)

type App struct {
	Out    io.Writer
	Config *config.Config
}

func New() (*App, error) {

	c, err := config.Load()
	if err != nil {
		return nil, err
	}

	a := &App{
		Out:    os.Stdout,
		Config: c,
	}
	return a, nil

}
