package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func HeaderMaxLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[header-max-length] value must be an integer")
		return
	}

	message, state = headerMaxLengthValidator(commit.Header, config.Level, config.Value.(int))
	return
}

func headerMaxLengthValidator(header string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Header can't be longer than %d characters", value)

	if len(header) == 0 {
		return
	}

	if len(header) > value {
		state = level
	}

	return
}
