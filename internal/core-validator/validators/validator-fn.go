package validators

import (
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

type ValidatorFn func(*parser.ConventionalCommit, validation.ValidationRuleConfig) (string, validation.ValidationState, error)
