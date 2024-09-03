package output

import (
	"fmt"
	"strings"
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

func groupWorkByDay(work []models.Work) map[string][]models.Work {
	grouped := make(map[string][]models.Work)
	for _, w := range work {
		date := helpers.GetHumanDate(w.Added)
		grouped[date] = append(grouped[date], w)
	}
	return grouped
}

func groupWorkByMonth(work []models.Work) map[string][]models.Work {
	grouped := make(map[string][]models.Work)
	for _, w := range work {
		month := w.Added.Month().String()
		grouped[month] = append(grouped[month], w)
	}
	return grouped
}

func groupWorkByYear(work []models.Work) map[string][]models.Work {
	grouped := make(map[string][]models.Work)
	for _, w := range work {
		year := fmt.Sprintf("%d", w.Added.Year())
		grouped[year] = append(grouped[year], w)
	}
	return grouped
}

func renderWorkItem(w models.Work) string {

	symbol := lipgloss.NewStyle().SetString("*").
		Foreground(tertiaryColor).
		PaddingRight(1).
		String()

	contentText := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
		Render(w.Content)

	return lipgloss.JoinHorizontal(0, symbol, " ", contentText)

}

func renderWorkItemGroup(groupName string, workItems []models.Work) string {
	stringBuilder := new(strings.Builder)

	// Print title
	{
		title := lipgloss.NewStyle().
			Margin(0, 0, 1, 0).
			Width(20).
			Align(lipgloss.Left).
			Foreground(primaryColor).
			Bold(true).
			Render(groupName)

		stringBuilder.WriteString(title + "\n")
	}

	// Print items
	for _, w := range workItems {
		stringBuilder.WriteString(renderWorkItem(w) + "\n")
	}
	return stringBuilder.String()
}
