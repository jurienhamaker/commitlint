package rules

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/charmbracelet/log"
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func References(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	values := []string{}

	log.Info(reflect.ValueOf(config.Value))
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
		err = errors.New("[references] value must be a string or an array/slice of strings")
		return
	}

	message, state = referencesValidator(commit.Raw, config.Level, config.Always, values)
	return
}

func referencesValidator(raw string, level validation.ValidationState, always bool, prefixes []string) (rule string, state validation.ValidationState) {
	rule = "Commit must contain reference(s)"
	if !always {
		rule = "Commit may not contain reference(s)"
	}

	if len(raw) == 0 {
		return
	}

	matches := []string{}
	for _, prefix := range prefixes {
		regex := regexp.MustCompile(fmt.Sprintf(`%s\w+`, prefix))
		if prefix == "sha" {
			regex = regexp.MustCompile(`(?i)\b[0-9a-f]{5,40}\b`)
		}

		match := regex.FindStringSubmatch(raw)
		matches = append(matches, match...)
	}

	if len(matches) == 0 && always {
		state = level
	}

	if len(matches) > 0 && !always {
		state = level
	}

	return
}
