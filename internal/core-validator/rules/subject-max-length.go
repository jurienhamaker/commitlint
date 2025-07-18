package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func SubjectMaxLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[subject-max-length] value must be an integer")
		return
	}

	message, state = subjectMaxLengthValidator(commit.Subject, config.Level, config.Value.(int))
	return
}

func subjectMaxLengthValidator(subject string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Subject can't be longer than %d characters", value)

	if len(subject) == 0 {
		return
	}

	if len(subject) > value {
		state = level
	}

	return
}
