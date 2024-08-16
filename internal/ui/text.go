package ui

import "github.com/charmbracelet/lipgloss"

var (
	TextErrorStyle = lipgloss.NewStyle().Bold(true).
			Foreground(errorC).
			PaddingTop(1).
			PaddingBottom(1)

	TextWarningStyle = lipgloss.NewStyle().Bold(true).
				Foreground(warning).
				PaddingTop(1).
				PaddingBottom(1)

	TextSuccessStyle = lipgloss.NewStyle().Bold(true).
				Foreground(success).
				PaddingTop(1).
				PaddingBottom(1)

	TextPrimaryStyle = lipgloss.NewStyle().Bold(true).
				Foreground(primary).
				PaddingTop(1).
				PaddingBottom(1)

	TextSecondaryStyle = lipgloss.NewStyle().Bold(true).
				Foreground(secondary).
				PaddingTop(1).
				PaddingBottom(1)

	TextTertiaryStyle = lipgloss.NewStyle().Bold(true).
				Foreground(tertiary).
				PaddingTop(1).
				PaddingBottom(1)
)

func TextError(text string) string {
	return TextErrorStyle.Render(text)
}

func TextWarning(text string) string {
	return TextWarningStyle.Render(text)
}

func TextSuccess(text string) string {
	return TextSuccessStyle.Render(text)
}

func TextPrimary(text string) string {
	return TextPrimaryStyle.Render(text)
}

func TextSecondary(text string) string {
	return TextSecondaryStyle.Render(text)
}

func TextTertiary(text string) string {
	return TextTertiaryStyle.Render(text)
}
