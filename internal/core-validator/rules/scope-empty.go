package rules

import (
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func ScopeEmpty(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	message, state = scopeEmptyValidator(commit.Scope, config.Level, config.Always)
	return
}

func scopeEmptyValidator(scope string, level validation.ValidationState, always bool) (rule string, state validation.ValidationState) {
	rule = "Scope must be empty"
	if !always {
		rule = "Scope may not be empty"
	}

	if len(scope) == 0 && !always {
		state = level
	}

	if len(scope) > 0 && always {
		state = level
	}

	return
}
