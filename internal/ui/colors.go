package ui

import "github.com/charmbracelet/lipgloss"

var (
	neutral = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}

	errorC  = lipgloss.AdaptiveColor{Light: "#EF4444", Dark: "#FF8383"}
	warning = lipgloss.AdaptiveColor{Light: "#F59E0B", Dark: "#FFD166"}
	success = lipgloss.AdaptiveColor{Light: "#10B981", Dark: "#6EE7B7"}

	primary   = lipgloss.AdaptiveColor{Light: "#3B82F6", Dark: "#93C5FD"}
	secondary = lipgloss.AdaptiveColor{Light: "#6366F1", Dark: "#A5B4FC"}
	tertiary  = lipgloss.AdaptiveColor{Light: "#8B5CF6", Dark: "#D6BCFA"}
)
