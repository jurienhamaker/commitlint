package utils

import (
	tea "github.com/charmbracelet/bubbletea"
)

func WaitForActivity[T any](sub chan T) tea.Cmd {
	return func() tea.Msg {
		return T(<-sub)
	}
}
