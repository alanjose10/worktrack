package output

import (
	"strings"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/models"
	"github.com/charmbracelet/lipgloss"
)

func BuildListTodoOutput(from time.Time, to time.Time, todos []models.Todo) string {
	stringBuilder := strings.Builder{}

	// Title
	{
		titleName := lipgloss.NewStyle().
			Margin(1, 0, 1, 0).Width(11).
			Align(lipgloss.Left).
			Foreground(primaryColor).
			Bold(true).
			Render("TODO ITEMS")

		titleDivider := lipgloss.NewStyle().
			Margin(1, 0, 1, 0).Width(3).
			Align(lipgloss.Center).
			Foreground(secondaryColor).
			Render("|")

		titleDate := lipgloss.NewStyle().
			Margin(1, 0, 1, 0).Width(40).
			Foreground(primaryColor).
			Render(helpers.GetHumanDate(from) + " - " + helpers.GetHumanDate(to))

		title := lipgloss.JoinHorizontal(0, titleName, titleDivider, titleDate)

		stringBuilder.WriteString(title + "\n")
	}

	// Body
	{
		if len(todos) > 0 {
			for _, item := range todos {
				stringBuilder.WriteString(renderTodoItem(item) + "\n")
			}
		}

	}

	return stringBuilder.String()
}

func BuildListBlockerOutput(from time.Time, to time.Time, todos []models.Blocker) string {
	return "Blockers between " + helpers.GetHumanDate(from) + " and " + helpers.GetHumanDate(to) + "\n"
}

func BuildListWorkOutput(from time.Time, to time.Time, todos []models.Work) string {
	return "Work done from " + helpers.GetHumanDate(from) + " to " + helpers.GetHumanDate(to) + "\n"
}
