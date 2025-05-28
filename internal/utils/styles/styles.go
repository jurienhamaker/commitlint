package styles

import "github.com/charmbracelet/lipgloss"

var (
	GrayishStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	ErrorStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#ff0000")).Foreground(lipgloss.Color("#ffffff")).Render
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render
)
