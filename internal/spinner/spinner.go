package spinner

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/internal/styles"
)

type (
	SpinnerErrorMsg         error
	SpinnerResultMsg[T any] struct {
		Result T
		Error  error
	}
)

type SpinnerModel[T any] struct {
	spinner  spinner.Model
	quitting bool
	message  string

	Error      error
	Result     SpinnerResultMsg[T]
	ResultChan chan SpinnerResultMsg[T]
}

func CreateSpinner[T any](message string) SpinnerModel[T] {
	s := spinner.New()
	s.Spinner = spinner.Moon
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return SpinnerModel[T]{spinner: s, ResultChan: make(chan SpinnerResultMsg[T]), message: message}
}

func (m SpinnerModel[any]) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, utils.WaitForActivity(m.ResultChan))
}

func (m SpinnerModel[any]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			m.Result = SpinnerResultMsg[any]{Error: errors.New("interrupted by user")}
			return m, tea.Quit
		default:
			return m, nil
		}

	case SpinnerResultMsg[any]:
		m.quitting = true
		m.Result = msg
		return m, tea.Quit

	case SpinnerErrorMsg:
		m.Error = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m SpinnerModel[any]) View() string {
	if m.Error != nil {
		return m.Error.Error()
	}

	str := fmt.Sprintf(
		"\n  %s %s...\n     %s\n",
		m.spinner.View(),
		m.message,
		styles.GrayishStyle("Press ctrl+c to cancel"),
	)

	if m.quitting {
		return ""
	}

	return str
}
