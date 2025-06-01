package validators

import (
	"strings"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func FooterLeadingBlank(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	// we check the body here because we want the last line of the body be an empty line
	message, state = footerLeadingBlankValidator(commit.Body, commit.Footer, config.Level, config.Always)
	return
}

func footerLeadingBlankValidator(body string, footer []string, level validation.ValidationState, always bool) (rule string, state validation.ValidationState) {
	rule = "Footer must have leading blank line"
	if !always {
		rule = "Footer may not have leading blank line"
	}

	if len(footer) == 0 {
		return
	}

	splitted := strings.Split(body, "\n")

	if splitted[len(splitted)-1] != "" && always {
		state = level
	}

	if splitted[len(splitted)-1] == "" && !always {
		state = level
	}

	return
}
