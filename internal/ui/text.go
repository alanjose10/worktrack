package ui

import "github.com/charmbracelet/lipgloss"

var (
	TextErrorStyle = lipgloss.NewStyle().Bold(true).
		Foreground(errorC).
		PaddingTop(1).
		PaddingBottom(1)
)

func TextError(text string) string {
	return TextErrorStyle.Render(text)
}
