package rules

import (
	"strings"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func SignedOffBy(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	message, state = signedOffByValidator(commit.Footer, config.Level, config.Always)
	return
}

func signedOffByValidator(footer []string, level validation.ValidationState, always bool) (rule string, state validation.ValidationState) {
	rule = "Footer must contain signed off by trailer"
	if !always {
		rule = "Footer may not contain signed off by trailer"
	}

	containsSignedOffBy := false
	for _, footerLine := range footer {
		if strings.HasPrefix(footerLine, "Signed-off-by:") {
			containsSignedOffBy = true
			break
		}
	}

	if (!containsSignedOffBy && always) || (containsSignedOffBy && !always) {
		state = level
	}

	return
}
