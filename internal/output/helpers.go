package output

import (
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/models"
	"github.com/charmbracelet/lipgloss"
)

var (
	doneTodoSymbol = lipgloss.
			NewStyle().
			SetString("[✓]").
			Foreground(tertiaryColor).
			PaddingRight(1).
			String()

	pendingTodoSymbol = lipgloss.NewStyle().SetString("[ ]").
				Foreground(tertiaryColor).
				PaddingRight(1).
				String()
)

func renderAddedOnDateByColor(d time.Time) string {
	if helpers.IsToday(d) {
		return lipgloss.NewStyle().
			Foreground(okColor).
			Render(" (Added today)")
	}
	if helpers.IsThisWeek(d) {
		return lipgloss.NewStyle().
			Foreground(primaryColor).
			Render(" (Added on " + helpers.GetHumanDate(d) + ")")
	}
	if helpers.IsThisMonth(d) {
		return lipgloss.NewStyle().
			Foreground(warningColor).
			Render(" (Added on " + helpers.GetHumanDate(d) + ")")
	}
	return lipgloss.NewStyle().
		Foreground(errorColor).
		Render(" (Added on " + helpers.GetHumanDate(d) + ")")

}

func renderTodoItem(t models.Todo) string {

	contentText := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
		Render(t.Content)

	if t.Done {
		return lipgloss.JoinHorizontal(0, doneTodoSymbol, contentText, renderAddedOnDateByColor(t.Added))
	} else {
		return lipgloss.JoinHorizontal(0, pendingTodoSymbol, contentText, renderAddedOnDateByColor(t.Added))
	}

}
