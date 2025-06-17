package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/jurienhamaker/commitlint/config"
	"github.com/jurienhamaker/commitlint/internal/spinner"
	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/internal/utils"
)

type Install struct {
	Global        bool `help:"Wether to install commitlint config globally"`
	RegisterHooks bool `help:"Wether to register commitlint global hooks to git"`
}

func (i Install) Run(ctx *kong.Context) error {
	message := "Installing commitlint in your repository"
	if i.Global {
		message = "Installing commitlint globally"
	}

	if i.RegisterHooks && !i.Global {
		utils.ReplyError("You can only register hooks to git if you install commitlint globally")
		os.Exit(0)
	}

	m := spinner.CreateSpinner[bool](message)
	p := tea.NewProgram(m)

	go install(m.ResultChan, i.Global, i.RegisterHooks)

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
		globalPath, _ := config.GetGlobalPath()
		if strings.Contains(globalPath, ".config") {
			homeDir, _ := config.GetUserHome()
			globalPath = strings.Replace(globalPath, homeDir, "~", 1)
		}

		if i.RegisterHooks {
			fmt.Printf(
				"\n%s\n\n",
				styles.SuccessTextStyle(
					"Success: Installed commitlint globally & registered hooks folder to git",
				),
			)
		} else {
			fmt.Printf(
				"\n%s\n\n%s\n%s\n%s\n\n",
				styles.SuccessTextStyle(
					"Success: Installed commitlint globally",
				),
				styles.WhiteTextStyle(
					"Want to globally install commit hooks?",
				),
				styles.LightGrayTextStyle(
					fmt.Sprintf(`Use "git config --global core.hooksPath %s/hooks"`, globalPath),
				),
				styles.SupportiveLilacTextStyleHyperlink(
					"Or check out our guide on how to add global hooks!", "https://commitlint.jurien.dev/guides/global-hooks",
				),
			)
		}
		return nil
	}

	utils.ReplySuccess("Installed commitlint in your repository")
	return nil
}
