package output

import "github.com/charmbracelet/lipgloss"

var (
	neutralColor = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}

	errorColor   = lipgloss.AdaptiveColor{Light: "#EF4444", Dark: "#FF8383"}
	warningColor = lipgloss.AdaptiveColor{Light: "#F59E0B", Dark: "#FFD166"}
	okColor      = lipgloss.AdaptiveColor{Light: "#10B981", Dark: "#6EE7B7"}

	primaryColor   = lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}
	secondaryColor = lipgloss.AdaptiveColor{Light: "#6366F1", Dark: "#A5B4FC"}
	tertiaryColor  = lipgloss.AdaptiveColor{Light: "#8B5CF6", Dark: "#D6BCFA"}
)
