package ui

import (
	"strings"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/models"
	"github.com/charmbracelet/lipgloss"
)

var (
	borderColor = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	border = lipgloss.Border{
		Top:    "─",
		Bottom: "─",
		Left:   "│",
		Right:  "│",
		// TopLeft:     "╭",
		// TopRight:    "╮",
		// BottomLeft:  "┴",
		// BottomRight: "┴",
	}

	standupReportStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2).Background(lipgloss.Color("#000000")).Border(border).BorderForeground(borderColor)

	titleStyle = lipgloss.NewStyle().
			Margin(1, 0, 0, 0).Width(100).
			Align(lipgloss.Center).
			Foreground(lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}).
			Bold(true)

	dateStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(borderColor).
			Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"})

	dailyUpdateHeader = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderBottom(true).
				BorderForeground(borderColor).
				MarginRight(2).
				Render
)

func PrintStandupReport(goBack int, todos []models.Todo, blockers []models.Blocker) string {
	// Print the standup report

	doc := strings.Builder{}

	today := helpers.GetCurrentDate()

	// Title
	{
		title := lipgloss.JoinVertical(
			lipgloss.Center,
			titleStyle.Render("STANDUP REPORT"),
			dateStyle.Render(today.Format("Monday, 02 January 2006")),
		)

		doc.WriteString(title + "\n\n")
	}

	{
		days := helpers.GetNPrevWorkingDays(today, goBack)

		for _, day := range days {
			doc.WriteString(dailyUpdateHeader(day.Format("Monday, 02 January 2006")) + "\n")

		}
	}

	return standupReportStyle.Render(doc.String())

}
