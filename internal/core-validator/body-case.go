package corevalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jurienhamaker/commitlint/internal/utils"
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func bodyCase(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	values := []string{}
	if reflect.ValueOf(config.Value).Kind() == reflect.String {
		values = append(values, config.Value.(string))
	}

	if reflect.ValueOf(config.Value).Kind() == reflect.Array {
		values = config.Value.([]string)
	}

	if len(values) == 0 {
		err = errors.New("[body-case] value must be a string or an array of strings")
		return
	}

	for _, str := range values {
		_, ok := utils.CaseStringToCase[str]
		if !ok {
			keys := []string{}
			for k := range utils.CaseStringToCase {
				keys = append(keys, k)
			}

			err = fmt.Errorf("[body-case] unknown value: %s. Must be any of: %s", str, strings.Join(keys, ", "))
			return
		}
	}

	message, state = bodyCaseValidator(commit.Body, config.Level, config.Always, values)
	return
}

func bodyCaseValidator(body string, level validation.ValidationState, always bool, cases []string) (rule string, state validation.ValidationState) {
	joined := strings.Join(cases, ",")
	rule = fmt.Sprintf("Body must be in %s", joined)
	if !always {
		rule = fmt.Sprintf("Body may not be in %s", joined)
	}

	if len(cases) > 1 {
		rule = fmt.Sprintf("Body must be in any of the following cases: %s", joined)
		if !always {
			rule = fmt.Sprintf("Body may not be in any of the following cases: %s", joined)
		}
	}

	if len(body) == 0 {
		return
	}

	results := make(map[string]bool)
	for _, str := range cases {
		parsedCase := utils.CaseStringToCase[str]
		results[str] = utils.EnsureCase(body, parsedCase)
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
