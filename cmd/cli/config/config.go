package config

import (
	"github.com/alanjose10/worktrack/internal/worktrack"
	"github.com/spf13/cobra"
)

func BuildConfigCommand(app *worktrack.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "config",
		Short: "Manage worktrack configuration",
		Long:  `Manage worktrack configuration.`,
	}
	command.AddCommand(buildInitCommand(app))
	command.AddCommand(buildViewCommand(app))

	return command
}
