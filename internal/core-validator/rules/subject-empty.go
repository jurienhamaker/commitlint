package rules

import (
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func SubjectEmpty(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	message, state = subjectEmptyValidator(commit.Subject, config.Level, config.Always)
	return
}

func subjectEmptyValidator(subject string, level validation.ValidationState, always bool) (rule string, state validation.ValidationState) {
	rule = "Subject must be empty"
	if !always {
		rule = "Subject may not be empty"
	}

	if len(subject) == 0 && !always {
		state = level
	}

	if len(subject) > 0 && always {
		state = level
	}

	return
}
