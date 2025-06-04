package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func HeaderFullStop(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.String {
		err = errors.New("[header-full-stop] value must be a string")
		return
	}

	message, state = headerFullStopValidator(commit.Header, config.Level, config.Always, config.Value.(string))

	return
}

func headerFullStopValidator(header string, level validation.ValidationState, always bool, value string) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Header must end with \"%s\"", value)
	if !always {
		rule = fmt.Sprintf("Header may not end with \"%s\"", value)
	}

	if len(header) == 0 {
		return
	}

	lastChar := header[len(header)-1:]
	if lastChar != value && always {
		state = level
		return
	}

	if lastChar == value && !always {
		state = level
		return
	}

	return
}
