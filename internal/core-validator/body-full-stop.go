package corevalidator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func bodyFullStop(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.String {
		err = errors.New("[body-full-stop] value must be a string")
		return
	}

	message, state = bodyFullStopValidator(commit.Body, config.Level, config.Always, config.Value.(string))

	return
}

func bodyFullStopValidator(body string, level validation.ValidationState, always bool, value string) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Body must end with \"%s\"", value)
	if !always {
		rule = fmt.Sprintf("Body may not end with \"%s\"", value)
	}

	if len(body) == 0 {
		return
	}

	lastChar := body[len(body)-1:]
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
