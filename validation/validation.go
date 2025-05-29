package validation

type (
	ValidationState   int
	ValidationResult  map[string]ValidationState
	ValidationsResult map[string]ValidationResult
	ValidatorConfig   map[string]any
	Validator         func(string, map[string]any) (ValidationResult, error)
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
