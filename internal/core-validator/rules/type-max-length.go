package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func TypeMaxLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[type-max-length] value must be an integer")
		return
	}

	message, state = typeMaxLengthValidator(commit.Type, config.Level, config.Value.(int))
	return
}

func typeMaxLengthValidator(typeStr string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Type can't be longer than %d characters", value)

	if len(typeStr) == 0 {
		return
	}

	if len(typeStr) > value {
		state = level
	}

	return
}
