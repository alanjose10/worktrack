package output

import (
	"sort"
	"strings"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/alanjose10/worktrack/internal/models"
	"github.com/charmbracelet/lipgloss"
)

func BuildListTodoOutput(from time.Time, to time.Time, todos []models.Todo) string {
	stringBuilder := new(strings.Builder)

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

func BuildListBlockerOutput(from time.Time, to time.Time, blockers []models.Blocker) string {
	stringBuilder := new(strings.Builder)

	// Title
	{
		titleName := lipgloss.NewStyle().
			Margin(1, 0, 1, 0).Width(10).
			Align(lipgloss.Left).
			Foreground(primaryColor).
			Bold(true).
			Render("BLOCKERS")

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

	sort.Slice(blockers, func(i, j int) bool {
		return blockers[i].Added.Before(blockers[j].Added)
	})

	// minDate := blockers[0].Added
	// maxDate := blockers[len(blockers)-1].Added

	// fmt.Printf("minDate: %v\n", minDate)
	// fmt.Printf("maxDate: %v\n", maxDate)

	// Body
	{
		if len(blockers) > 0 {
			for _, item := range blockers {
				stringBuilder.WriteString(renderBlocker(item) + "\n")
			}
		}
	}

	return stringBuilder.String()
}

func BuildListWorkOutput(from time.Time, to time.Time, workItems []models.Work) string {
	stringBuilder := new(strings.Builder)

	// Title
	{
		titleName := lipgloss.NewStyle().
			Margin(1, 0, 1, 0).Width(11).
			Align(lipgloss.Left).
			Foreground(primaryColor).
			Bold(true).
			Render("WORK ITEMS")

		titleDivider := lipgloss.NewStyle().
			Margin(1, 0, 1, 0).Width(3).
			Align(lipgloss.Center).
			Foreground(secondaryColor).
			Bold(true).
			Render("|")

		titleDate := lipgloss.NewStyle().
			Margin(1, 0, 1, 0).Width(40).
			Foreground(primaryColor).
			Render(helpers.GetHumanDate(from) + " - " + helpers.GetHumanDate(to))

		title := lipgloss.JoinHorizontal(0, titleName, titleDivider, titleDate)

		stringBuilder.WriteString(title + "\n")
	}

	sort.Slice(workItems, func(i, j int) bool {
		return workItems[i].Added.Before(workItems[j].Added)
	})

	minDate := workItems[0].Added
	maxDate := workItems[len(workItems)-1].Added

	// Body
	{
		if len(workItems) == 0 {
			stringBuilder.WriteString("No work items found\n")
			return stringBuilder.String()
		}

		var groupdWorkItems map[string][]models.Work
		if maxDate.Sub(minDate) > time.Hour*24*365 {
			groupdWorkItems = groupWorkByYear(workItems)
		} else if maxDate.Sub(minDate) > time.Hour*24*30 {
			groupdWorkItems = groupWorkByMonth(workItems)
		} else if maxDate.Sub(minDate) > 0 {
			groupdWorkItems = groupWorkByDay(workItems)
		}

		for key, group := range groupdWorkItems {
			stringBuilder.WriteString(renderWorkItemGroup(key, group) + "\n")
		}

	}

	return stringBuilder.String()
}
