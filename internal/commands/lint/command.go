package lint

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/jurienhamaker/commitlint/config"
	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/validation"
)

type Lint struct {
	Message []string `arg:"" help:"The message to lint" default:""`
}

func (i Lint) Run(ctx *kong.Context) error {
	err := config.Load()
	if err != nil {
		utils.ReplyError(err.Error())
		os.Exit(1)
	}

	config := config.GetConfig()
	if !config.Enabled {
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
	if isInputPiped() {
		data, dataErr := io.ReadAll(os.Stdin)
		if dataErr != nil {
			err = dataErr
			return
		}

		dataStr := string(data)
		dataStr = strings.TrimRight(dataStr, "\n")
		input = []string{dataStr}
	}

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

func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
