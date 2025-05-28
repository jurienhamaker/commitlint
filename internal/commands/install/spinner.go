package install

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/internal/utils/styles"
)

type (
	errMsg    error
	resultMsg struct {
		Installed bool
		Error     error
	}
)

type model struct {
	spinner    spinner.Model
	quitting   bool
	err        error
	result     resultMsg
	resultChan chan resultMsg
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Moon
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{spinner: s, resultChan: make(chan resultMsg)}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, utils.WaitForActivity(m.resultChan))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}

	case resultMsg:
		m.quitting = true
		m.result = msg
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	str := fmt.Sprintf(
		"\n\n   %s Installing commitlint on the repository...\n      %s\n\n",
		m.spinner.View(),
		styles.GrayishStyle("Press ctrl+c to cancel"),
	)

	if m.quitting {
		return "\n"
	}

	return str
}
