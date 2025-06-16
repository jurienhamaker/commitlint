package lint

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/jurienhamaker/commitlint/internal/spinner"
	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/plugins"
	"github.com/jurienhamaker/commitlint/validation"
)

func runPlugins(pm *plugins.PluginManager, message string) (validation.ValidationsResult, error) {
	m := spinner.CreateSpinner[validation.ValidationsResult]("Linting message")
	p := tea.NewProgram(m)

	go func(sub chan spinner.SpinnerResultMsg[validation.ValidationsResult]) {
		commit := parser.ParseConventionalCommit(message)
		result, err := pm.RunPluginValidators(commit)
		if err != nil {
			sub <- spinner.SpinnerResultMsg[validation.ValidationsResult]{Error: fmt.Errorf("running failed: %s", err.Error())}
			return
		}

		sub <- spinner.SpinnerResultMsg[validation.ValidationsResult]{Result: result}
	}(m.ResultChan)

	run, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run program: %s", err)
	}

	result := run.(spinner.SpinnerModel[validation.ValidationsResult]).Result
	if result.Error != nil {
		return nil, result.Error
	}

	log.Debug(styles.SuccessTextStyle("Success: Ran plugin validators"))
	return result.Result, nil
}
