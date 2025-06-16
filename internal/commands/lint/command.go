package lint

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/jurienhamaker/commitlint/config"
	corevalidator "github.com/jurienhamaker/commitlint/internal/core-validator"
	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/plugins"
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

	var result validation.ValidationsResult
	if runtime.GOOS == "windows" {
		result, err = windoosRun(message)
	} else {
		result, err = unixRun(message)
	}

	if err != nil {
		utils.ReplyError(err.Error())
		os.Exit(1)
	}

	total, parseResult := parseResult(message, result)

	fmt.Println(
		styles.GrayishStyle(
			fmt.Sprintf(
				"%d rules checked. %d success, %d warnings & %d errors\n",
				total,
				parseResult[validation.ValidationStateSuccess],
				parseResult[validation.ValidationStateWarning],
				parseResult[validation.ValidationStateError],
			),
		),
	)
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

func unixRun(message string) (validation.ValidationsResult, error) {
	pm, err := loadPlugins()
	if err != nil {
		return nil, fmt.Errorf("unix run error: %s", err)
	}

	result, err := runPlugins(pm, message)
	if err != nil {
		return nil, fmt.Errorf("unix run error: %s", err)
	}

	return result, nil
}

func windoosRun(message string) (validation.ValidationsResult, error) {
	pm, err := plugins.LoadPlugins(".commitlint/plugins")
	if err != nil {
		return nil, fmt.Errorf("windows run error: %s", err)
	}
	pm.RegisterPlugin("core", corevalidator.CoreValidator)

	commit := parser.ParseConventionalCommit(message)
	result, err := pm.RunPluginValidators(commit)
	if err != nil {
		return nil, fmt.Errorf("windows run error: %s", err)
	}

	return result, nil
}

func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
