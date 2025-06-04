package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func FooterMaxLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[footer-max-length] value must be an integer")
		return
	}

	message, state = footerMaxLengthValidator(commit.Footer, config.Level, config.Value.(int))
	return
}

func footerMaxLengthValidator(footer []string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Footer can't be longer than %d characters", value)

	if len(footer) == 0 {
		return
	}

	// join so we can count total length
	joined := strings.Join(footer, "\n")
	if len(joined) > value {
		state = level
	}

	return
}
