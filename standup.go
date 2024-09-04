package main

import (
	"os"
	"strings"

	"github.com/alanjose10/worktrack/internal/helpers"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	borderColor = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	border = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
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
				Foreground(lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}).
				Render

	workItemSymbol = lipgloss.NewStyle().SetString("->").
			Foreground(lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}).
			PaddingRight(1).
			String()

	doneTodoSymbol = lipgloss.NewStyle().SetString("[✓]").
			Foreground(lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}).
			PaddingRight(1).
			String()

	pendingTodoSymbol = lipgloss.NewStyle().SetString("[ ]").
				Foreground(lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}).
				PaddingRight(1).
				String()

	blockerSymbol = lipgloss.NewStyle().SetString("X").
			Foreground(lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}).
			PaddingRight(1).
			String()
)

func renderPendingTodoItem(s string, width int) string {
	return pendingTodoSymbol + lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#92400e", Dark: "#fcd34d"}).
		Width(width).
		Render(s)
}

func renderTodoItem(s string, width int) string {
	return doneTodoSymbol + lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
		Width(width).
		Render(s)
}

func renderWorkItem(s string, width int) string {
	return workItemSymbol + lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
		Width(width).
		Render(s)
}

func renderBlockerItem(s string, width int) string {
	return blockerSymbol + lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#7f1d1d", Dark: "#f87171"}).
		Width(width).
		Render(s)
}

func renderResolvedBlockerItem(s string, width int) string {
	return blockerSymbol + lipgloss.NewStyle().
		Strikethrough(true).
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
		Width(width).
		Render(s)
}

func (app *application) printStandupReport(goBack int) string {

	screenWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
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

			from := helpers.GetDateFloor(day)
			to := helpers.GetDateCeil(day)

			workList, err := app.workModel.ListBetween(from, to)
			if err != nil {
				os.Exit(1)
			}

			if len(workList) == 0 {
				continue
			}

			doc.WriteString(dailyUpdateHeader(day.Format("Monday, 02 January 2006")) + "\n")

			for _, work := range workList {
				doc.WriteString(renderWorkItem(work.Content, screenWidth-10) + "\n")
			}

			doc.WriteString("\n\n")

		}
	}

	{

		todos, err := app.todoModel.List()
		if err != nil {
			os.Exit(1)
		}

		if len(todos) > 0 {
			doc.WriteString(dailyUpdateHeader("TODO") + "\n")

			for _, item := range todos {
				if item.Done {
					doc.WriteString(renderTodoItem(item.Content, screenWidth-10) + "\n")
				} else {
					doc.WriteString(renderPendingTodoItem(item.Content, screenWidth-10) + "\n")
				}
			}

			doc.WriteString("\n\n")
		}

	}

	{

		blockers, err := app.blockerModel.List()
		if err != nil {
			os.Exit(1)
		}

		if len(blockers) > 0 {
			doc.WriteString(dailyUpdateHeader("BLOCKERS") + "\n")

			for _, item := range blockers {
				if item.Resolved {
					doc.WriteString(renderResolvedBlockerItem(item.Content, screenWidth-10) + "\n")
				} else {
					doc.WriteString(renderBlockerItem(item.Content, screenWidth-10) + "\n")
				}
			}

		}

	}

	doc.WriteString("\n")

	return standupReportStyle.Render(doc.String())

}
