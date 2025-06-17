package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/jurienhamaker/commitlint/internal/spinner"
	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/internal/utils"
)

type Install struct {
	Global bool `help:"Wether to install commitlint config globally"`
}

func (i Install) Run(ctx *kong.Context) error {
	message := "Installing commitlint in your repository"
	if i.Global {
		message = "Installing commitlint globally"
	}

	m := spinner.CreateSpinner[bool](message)
	p := tea.NewProgram(m)

	go install(m.ResultChan, i.Global)

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

	if i.Global {
		fmt.Printf(
			"\n%s\n%s %s\n\n",
			styles.SuccessTextStyle(
				"Success: Installed commitlint globally",
			),
			styles.GrayishTextStyle(
				"Want to globally install commit hooks?",
			),
			styles.SupportiveLilacTextStyleHyperlink(
				"Click here to check out our guide!", "https://commitlint.jurien.dev/guides/global-hooks",
			),
		)
		return nil
	}

	utils.ReplySuccess("Installed commitlint in your repository")
	return nil
}
