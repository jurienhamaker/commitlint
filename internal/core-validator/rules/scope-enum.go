package rules

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func ScopeEnum(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
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
		err = errors.New("[scope-case] value must be a string or an array/slice of strings")
		return
	}

	message, state = scopeEnumValidator(commit.Scope, config.Level, config.Always, values)
	return
}

func scopeEnumValidator(scope string, level validation.ValidationState, always bool, values []string) (rule string, state validation.ValidationState) {
	joined := strings.Join(values, ",")
	rule = fmt.Sprintf("Scope must be %s", joined)
	if !always {
		rule = fmt.Sprintf("Scope may not be %s", joined)
	}

	if len(values) > 1 {
		rule = fmt.Sprintf("Scope must be any of the following: %s", joined)
		if !always {
			rule = fmt.Sprintf("Scope may not be any of the following: %s", joined)
		}
	}

	if len(scope) == 0 {
		return
	}

	re := regexp.MustCompile(`/|\\|, ?`)
	split := re.Split(scope, -1)

	results := make(map[string]bool)
	for _, value := range values {
		ensured := slices.Contains(split, value)
		results[value] = ensured
	}

	failures := []string{}
	for str, result := range results {
		if (!result && always) || (result && !always) {
			failures = append(failures, str)
		}
	}

	if len(failures) > 0 {
		state = level
		return
	}

	return
}
