package install

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jurienhamaker/commitlint/internal/utils/styles"
)

type Install struct{}

func (i Install) Run(ctx *kong.Context) error {
	m := initialModel()
	p := tea.NewProgram(m)

	go install(m.resultChan)

	run, err := p.Run()
	if err != nil {
		fmt.Println(styles.ErrorStyle("Error: "+err.Error()) + "\n")
		os.Exit(1)
	}

	result := run.(model).result
	if result.Error != nil {
		fmt.Println(styles.ErrorStyle("Error: "+result.Error.Error()) + "\n")
		os.Exit(1)
	}

	fmt.Println(styles.SuccessStyle("Success: Installed commitlint in your repository") + "\n")
	return nil
}
