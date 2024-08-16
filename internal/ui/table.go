package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	HeaderStyle = lipgloss.NewStyle().Padding(0, 1, 0, 1).Foreground(neutral).Bold(true).Align(lipgloss.Center)
	CellStyle   = lipgloss.NewStyle().Padding(0, 1, 0, 1)
)

func TasksTable(columns []string, rows [][]string) string {

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(secondary)).
		StyleFunc(func(row, col int) lipgloss.Style {

			if row == 0 {
				return HeaderStyle
			} else {
				if columns[col] == "Status" {
					status := rows[row-1][col]
					switch status {
					case "done":
						return CellStyle.Foreground(success)
					case "in progress":
						return CellStyle.Foreground(primary)
					case "todo":
						return CellStyle.Foreground(warning)
					case "blocked":
						return CellStyle.Foreground(errorC)
					}
				}
				return CellStyle.Foreground(primary)
			}
		}).
		Headers(columns...).
		Rows(rows...)
	return t.Render()
}
