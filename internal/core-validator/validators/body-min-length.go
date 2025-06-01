package validators

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func BodyMinLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[body-min-length] value must be an integer")
		return
	}

	message, state = bodyMinLengthValidator(commit.Body, config.Level, config.Value.(int))
	return
}

func bodyMinLengthValidator(body string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Body must be atleast %d characters", value)

	if len(body) == 0 {
		return
	}

	if len(body) < value {
		state = level
	}

	return
}
