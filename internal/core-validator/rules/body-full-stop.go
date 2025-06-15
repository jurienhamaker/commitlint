package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func BodyFullStop(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	value := config.Value
	parsedValue := "."

	if value != nil && reflect.ValueOf(value).Kind() != reflect.String {
		err = errors.New("[body-full-stop] value must be a string")
		return
	}

	if value != nil {
		parsedValue = value.(string)
	}

	message, state = bodyFullStopValidator(commit.Body, config.Level, config.Always, parsedValue)
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
