package lint

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	corevalidator "github.com/jurienhamaker/commitlint/internal/core-validator"
	"github.com/jurienhamaker/commitlint/internal/spinner"
	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/plugins"
)

func loadPlugins() (*plugins.PluginManager, error) {
	m := spinner.CreateSpinner[*plugins.PluginManager]("Loading plugins")
	p := tea.NewProgram(m)

	go func(sub chan spinner.SpinnerResultMsg[*plugins.PluginManager]) {
		pm, err := plugins.LoadPlugins(".commitlint/plugins")
		if err != nil {
			sub <- spinner.SpinnerResultMsg[*plugins.PluginManager]{Error: fmt.Errorf("could not load plugins: %s", err.Error())}
			return
		}

		pm.RegisterPlugin("core", corevalidator.CoreValidator)

		sub <- spinner.SpinnerResultMsg[*plugins.PluginManager]{Result: pm}
	}(m.ResultChan)

	run, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run program: %s", err)
	}

	result := run.(spinner.SpinnerModel[*plugins.PluginManager]).Result
	if result.Error != nil {
		return nil, result.Error
	}

	log.Debug(styles.SuccessTextStyle("Success: Loaded plugins"))
	return result.Result, nil
}
