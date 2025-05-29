package validation

type ValidationRuleConfig struct {
	Enabled bool
	Level   ValidationState
	Value   any
}

type (
	ValidationState   int
	ValidationResult  map[string]ValidationState
	ValidationsResult map[string]ValidationResult
	ValidatorConfig   map[string]ValidationRuleConfig
	Validator         func(string, ValidatorConfig) (ValidationResult, error)
)

const (
	ValidationStateError ValidationState = iota
	ValidationStateWarning
	ValidationStateSuccess
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
