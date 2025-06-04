package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func FooterMinLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[footer-min-length] value must be an integer")
		return
	}

	message, state = footerMinLengthValidator(commit.Footer, config.Level, config.Value.(int))
	return
}

func footerMinLengthValidator(footer []string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Footer must be atleast %d characters", value)

	if len(footer) == 0 {
		return
	}

	// join so we can count total length
	joined := strings.Join(footer, "\n")
	if len(joined) < value {
		state = level
	}

	return
}
