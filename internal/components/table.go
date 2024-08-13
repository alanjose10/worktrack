package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	HeaderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true).Align(lipgloss.Center)
	CellStyle   = lipgloss.NewStyle().Padding(0, 1)
)

func Table(columns []string, rows [][]string) string {

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(blue)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			default:
				return CellStyle
			}
		}).
		Headers(columns...).
		Rows(rows...)
	return t.Render()
}
