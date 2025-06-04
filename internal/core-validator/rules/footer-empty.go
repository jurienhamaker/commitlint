package rules

import (
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func FooterEmpty(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	message, state = footerEmptyValidator(commit.Footer, config.Level, config.Always)
	return
}

func footerEmptyValidator(footer []string, level validation.ValidationState, always bool) (rule string, state validation.ValidationState) {
	rule = "Footer must be empty"
	if !always {
		rule = "Footer may not be empty"
	}

	if len(footer) == 0 && !always {
		state = level
	}

	if len(footer) > 0 && always {
		state = level
	}

	return
}
