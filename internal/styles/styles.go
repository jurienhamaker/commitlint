package styles

import "github.com/charmbracelet/lipgloss"

var (
	BoldTextStyle = lipgloss.NewStyle().Bold(true).Render

	ErrorStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#cc3300")).Foreground(lipgloss.Color("#ffffff")).Render
	WarningStyle = lipgloss.NewStyle().Background(lipgloss.Color("#ff9966")).Foreground(lipgloss.Color("#ffffff")).Render
	SuccessStyle = lipgloss.NewStyle().Background(lipgloss.Color("#99cc33")).Foreground(lipgloss.Color("#ffffff")).Render

	GrayishTextStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	SupportiveLilacTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D1AEFF")).Render
	ErrorTextStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#cc3300")).Render
	WarningTextStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff9966")).Render
	SuccessTextStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#99cc33")).Render
)
