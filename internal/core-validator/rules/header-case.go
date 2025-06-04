package rules

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func HeaderCase(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
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
		err = errors.New("[header-case] value must be a string or an array/slice of strings")
		return
	}

	for _, str := range values {
		_, ok := utils.CaseStringToCase[str]
		if !ok {
			keys := []string{}
			for k := range utils.CaseStringToCase {
				keys = append(keys, k)
			}

			err = fmt.Errorf("[header-case] unknown value: %s. Must be any of: %s", str, strings.Join(keys, ", "))
			return
		}
	}

	message, state = headerCaseValidator(commit.Header, config.Level, config.Always, values)
	return
}

func headerCaseValidator(header string, level validation.ValidationState, always bool, cases []string) (rule string, state validation.ValidationState) {
	joined := strings.Join(cases, ",")
	rule = fmt.Sprintf("Header must be in %s", joined)
	if !always {
		rule = fmt.Sprintf("Header may not be in %s", joined)
	}

	if len(cases) > 1 {
		rule = fmt.Sprintf("Header must be in any of the following cases: %s", joined)
		if !always {
			rule = fmt.Sprintf("Header may not be in any of the following cases: %s", joined)
		}
	}

	regex := regexp.MustCompile(`(?i)[a-z]`)
	match := regex.FindStringSubmatch(header)
	if len(header) == 0 || len(match) == 0 {
		return
	}

	results := make(map[string]bool)
	for _, str := range cases {
		parsedCase := utils.CaseStringToCase[str]
		results[str] = utils.EnsureCase(header, parsedCase)
	}

	failures := []string{}
	for str, result := range results {
		if (!result && always) || (result && !always) {
			failures = append(failures, str)
		}
	}

	if (len(failures) == len(cases) && always) || (len(failures) > 0 && !always) {
		state = level
		return
	}

	return
}
