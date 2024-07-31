package main

import (
	"log"

	"github.com/alanjose10/worktrack/cmd/cli/root"
	"github.com/alanjose10/worktrack/internal/worktrack"
)

func main() {
	app, err := worktrack.New()
	if err != nil {
		log.Fatal(err)
	}

	cmd := root.BuildRootCommand(app)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
