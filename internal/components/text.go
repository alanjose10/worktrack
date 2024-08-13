package components

import "github.com/charmbracelet/lipgloss"

var (
	TextErrorStyle = lipgloss.NewStyle().Bold(true).
		Foreground(red).
		PaddingTop(1).
		PaddingBottom(1)
)

func TextError(text string) string {
	return TextErrorStyle.Render(text)
}
