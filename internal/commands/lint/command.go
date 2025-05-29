package lint

import (
	"errors"
	"os"

	"github.com/alecthomas/kong"
	"github.com/jurienhamaker/commitlint/config"
	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/validation"
)

type Lint struct {
	Message []string `arg:"" help:"The message to lint" default:""`
}

func (i Lint) Run(ctx *kong.Context) error {
	config, err := config.Load()
	if err != nil {
		utils.ReplyError(err.Error())
		os.Exit(1)
	}

	if !config.Viper.GetBool("enabled") {
		utils.ReplyWarning("Commitlint is disabled")
		os.Exit(0)
	}

	message, err := validateInput(i.Message)
	if err != nil {
		utils.ReplyError(err.Error())
		os.Exit(1)
	}

	pm, err := loadPlugins()
	if err != nil {
		utils.ReplyError(err.Error())
		os.Exit(1)
	}

	result, err := runPlugins(pm, message)
	if err != nil {
		utils.ReplyError(err.Error())
		os.Exit(1)
	}

	parseResult := parseResult(result, message)
	if parseResult[validation.ValidationStateError] > 0 {
		os.Exit(1)
	}

	return nil
}

func validateInput(input []string) (message string, err error) {
	if len(input) == 0 {
		err = errors.New("no message given")
		return
	}

	if len(input) > 1 {
		for _, msg := range input {
			message += msg + " "
		}
	} else {
		message = input[0]
	}

	return
}
