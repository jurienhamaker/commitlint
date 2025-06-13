package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func ScopeMaxLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[scope-max-length] value must be an integer")
		return
	}

	message, state = scopeMaxLengthValidator(commit.Scope, config.Level, config.Value.(int))
	return
}

func scopeMaxLengthValidator(scope string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Scope can't be longer than %d characters", value)

	if len(scope) == 0 {
		return
	}

	if len(scope) > value {
		state = level
	}

	return
}
