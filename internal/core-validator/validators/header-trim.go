package validators

import (
	"strings"
	"unicode"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func HeaderTrim(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	message, state = headerTrimValidator(commit.Header, config.Level)
	return
}

func headerTrimValidator(header string, level validation.ValidationState) (rule string, state validation.ValidationState) {
	rule = "Header must not be surrounded by whitespaces"

	if len(header) == 0 {
		return
	}

	startsWithSpace := len(header) > len(strings.TrimLeftFunc(header, unicode.IsSpace))
	endsWithSpace := len(header) > len(strings.TrimRightFunc(header, unicode.IsSpace))

	if startsWithSpace && endsWithSpace {
		state = level
		return
	}

	if startsWithSpace {
		rule = "Header must not start with whitespace"
		state = level
	}

	if endsWithSpace {
		rule = "Header must not end with whitespace"
		state = level
	}

	return
}
