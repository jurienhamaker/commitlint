package rules

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func ScopeCase(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	values := []string{}
	if reflect.ValueOf(config.Value).Kind() == reflect.String {
		values = append(values, config.Value.(string))
	}

	if reflect.ValueOf(config.Value).Kind() == reflect.Slice || reflect.ValueOf(config.Value).Kind() == reflect.Array {
		parsed := reflect.ValueOf(config.Value)
		for i := range parsed.Len() {
			values = append(values, parsed.Index(i).Interface().(string))
		}
	}

	if len(values) == 0 {
		err = errors.New("[scope-case] value must be a string or an array/slice of strings")
		return
	}

	for _, str := range values {
		_, ok := utils.CaseStringToCase[str]
		if !ok {
			keys := []string{}
			for k := range utils.CaseStringToCase {
				keys = append(keys, k)
			}

			err = fmt.Errorf("[scope-case] unknown value: %s. Must be any of: %s", str, strings.Join(keys, ", "))
			return
		}
	}

	message, state = scopeCaseValidator(commit.Scope, config.Level, config.Always, values)
	return
}

func scopeCaseValidator(scope string, level validation.ValidationState, always bool, cases []string) (rule string, state validation.ValidationState) {
	joined := strings.Join(cases, ",")
	rule = fmt.Sprintf("Scope must be in %s", joined)
	if !always {
		rule = fmt.Sprintf("Scope may not be in %s", joined)
	}

	if len(cases) > 1 {
		rule = fmt.Sprintf("Scope must be in any of the following cases: %s", joined)
		if !always {
			rule = fmt.Sprintf("Scope may not be in any of the following cases: %s", joined)
		}
	}

	if len(scope) == 0 {
		return
	}

	re := regexp.MustCompile(`/|\\|, ?`)
	split := re.Split(scope, -1)

	results := make(map[string]bool)
	for _, s := range split {
		localResults := []bool{}
		for _, str := range cases {
			parsedCase := utils.CaseStringToCase[str]
			ensured := utils.EnsureCase(s, parsedCase)
			localResults = append(localResults, ensured)
		}

		results[s] = slices.Contains(localResults, true)
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
