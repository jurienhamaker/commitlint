package rules

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func TypeEnum(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	values := []string{}
	if reflect.ValueOf(config.Value).Kind() == reflect.String {
		values = append(values, config.Value.(string))
	}

	if reflect.ValueOf(config.Value).Kind() == reflect.Slice || reflect.ValueOf(config.Value).Kind() == reflect.Array {
		parsed := reflect.ValueOf(config.Value)
		for i := range parsed.Len() {
			values = append(values, parsed.Index(i).String())
		}
	}

	if len(values) == 0 {
		err = errors.New("[type-case] value must be a string or an array/slice of strings")
		return
	}

	message, state = typeEnumValidator(commit.Type, config.Level, config.Always, values)
	return
}

func typeEnumValidator(typeStr string, level validation.ValidationState, always bool, values []string) (rule string, state validation.ValidationState) {
	joined := strings.Join(values, ",")
	rule = fmt.Sprintf("Type must be %s", joined)
	if !always {
		rule = fmt.Sprintf("Type may not be %s", joined)
	}

	if len(values) > 1 {
		rule = fmt.Sprintf("Type must be any of the following: %s", joined)
		if !always {
			rule = fmt.Sprintf("Type may not be any of the following: %s", joined)
		}
	}

	if len(typeStr) == 0 {
		return
	}

	contains := slices.Contains(values, typeStr)

	if (!contains && always) || (contains && !always) {
		state = level
		return
	}

	return
}
