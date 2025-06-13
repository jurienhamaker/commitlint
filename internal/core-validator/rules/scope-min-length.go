package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func ScopeMinLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[scope-min-length] value must be an integer")
		return
	}

	message, state = scopeMinLengthValidator(commit.Scope, config.Level, config.Value.(int))
	return
}

func scopeMinLengthValidator(scope string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Scope must be atleast %d characters", value)

	if len(scope) == 0 {
		return
	}

	if len(scope) < value {
		state = level
	}

	return
}
