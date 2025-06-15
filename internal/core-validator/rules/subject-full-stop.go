package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func SubjectFullStop(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	value := config.Value
	parsedValue := "."

	if value != nil && reflect.ValueOf(value).Kind() != reflect.String {
		err = errors.New("[body-full-stop] value must be a string")
		return
	}

	if value != nil {
		parsedValue = value.(string)
	}

	message, state = subjectFullStopValidator(commit.Body, config.Level, config.Always, parsedValue)
	return
}

func subjectFullStopValidator(subject string, level validation.ValidationState, always bool, value string) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Subject must end with \"%s\"", value)
	if !always {
		rule = fmt.Sprintf("Subject may not end with \"%s\"", value)
	}

	if len(subject) == 0 {
		return
	}

	lastChar := subject[len(subject)-1:]
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
