package rules

import (
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func TypeEmpty(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	message, state = typeEmptyValidator(commit.Type, config.Level, config.Always)
	return
}

func typeEmptyValidator(typeStr string, level validation.ValidationState, always bool) (rule string, state validation.ValidationState) {
	rule = "Type must be empty"
	if !always {
		rule = "Type may not be empty"
	}

	if len(typeStr) == 0 && !always {
		state = level
	}

	if len(typeStr) > 0 && always {
		state = level
	}

	return
}
