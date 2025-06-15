package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func TypeMinLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[type-min-length] value must be an integer")
		return
	}

	message, state = typeMinLengthValidator(commit.Type, config.Level, config.Value.(int))
	return
}

func typeMinLengthValidator(typeStr string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Type must be atleast %d characters", value)

	if len(typeStr) == 0 {
		return
	}

	if len(typeStr) < value {
		state = level
	}

	return
}
