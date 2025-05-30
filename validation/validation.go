package validation

import "github.com/jurienhamaker/commitlint/parser"

type ValidationRuleConfig struct {
	Always bool
	Level  ValidationState
	Value  any
}

type (
	ValidationState   int
	ValidationResult  map[string]ValidationState
	ValidationsResult []ValidationResult
	ValidatorConfig   map[string]ValidationRuleConfig
	Validator         func(*parser.ConventionalCommit, ValidatorConfig) (ValidationResult, error)
)

const (
	ValidationStateSuccess ValidationState = iota
	ValidationStateWarning
	ValidationStateError
)

var ValidationStateName = map[ValidationState]string{
	ValidationStateError:   "error",
	ValidationStateWarning: "warning",
	ValidationStateSuccess: "success",
}

var ValidationStateInt = map[string]ValidationState{
	"error":   ValidationStateError,
	"warning": ValidationStateWarning,
	"success": ValidationStateSuccess,
}

var ValidationStateMapping = map[int]ValidationState{
	0: ValidationStateSuccess,
	1: ValidationStateWarning,
	2: ValidationStateError,
}
