package install

import (
	"os"
	"strings"

	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/jurienhamaker/commitlint/internal/spinner"
	"github.com/jurienhamaker/commitlint/internal/utils"
)

type Install struct{}

func (i Install) Run(ctx *kong.Context) error {
	m := spinner.CreateSpinner[bool]("Installing commitlint in your repository")
	p := tea.NewProgram(m)

	go install(m.ResultChan)

	run, err := p.Run()
	if err != nil {
		utils.ReplyError(err.Error())
		os.Exit(1)
	}

	result := run.(spinner.SpinnerModel[bool]).Result
	if result.Error != nil {
		if !strings.Contains(result.Error.Error(), "already exists") {
			utils.ReplyError(result.Error.Error())
			os.Exit(1)
		}

		utils.ReplyWarning(result.Error.Error())
	}

	utils.ReplySuccess("Installed commitlint in your repository")
	return nil
}
