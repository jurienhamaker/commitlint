package validators

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func FooterMaxLineLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[footer-max-line-length] value must be an integer")
		return
	}

	message, state = footerMaxLineLengthValidator(commit.Footer, config.Level, config.Value.(int))
	return
}

func footerMaxLineLengthValidator(footer []string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Footer can't have lines longer than %d characters", value)

	if len(footer) == 0 {
		return
	}

	for _, line := range footer {
		if len(line) > value {
			state = level
			break
		}
	}

	return
}
