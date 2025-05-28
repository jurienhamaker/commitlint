package lint

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/jurienhamaker/commitlint/internal/utils/styles"
)

type Lint struct {
	Message []string `arg:"" help:"The message to lint" default:""`
}

func (i Lint) Run(ctx *kong.Context) error {
	input := i.Message

	if len(input) == 0 {
		log.Info(styles.ErrorStyle("No message given"))
		os.Exit(1)
	}

	message := ""
	if len(input) > 1 {
		for _, msg := range input {
			message += msg + " "
		}
	} else {
		message = input[0]
	}

	log.Info(styles.SuccessStyle("Linting message: " + message))

	os.Exit(1)
	return nil
}
