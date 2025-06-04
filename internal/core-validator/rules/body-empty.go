package rules

import (
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func BodyEmpty(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	message, state = bodyEmptyValidator(commit.Body, config.Level, config.Always)
	return
}

func bodyEmptyValidator(body string, level validation.ValidationState, always bool) (rule string, state validation.ValidationState) {
	rule = "Body must be empty"
	if !always {
		rule = "Body may not be empty"
	}

	if len(body) == 0 && !always {
		state = level
	}

	if len(body) > 0 && always {
		state = level
	}

	return
}
