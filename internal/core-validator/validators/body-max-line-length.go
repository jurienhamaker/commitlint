package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func BodyMaxLineLength(commit *parser.ConventionalCommit, config validation.ValidationRuleConfig) (message string, state validation.ValidationState, err error) {
	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		err = errors.New("[body-max-line-length] value must be an integer")
		return
	}

	message, state = bodyMaxLineLengthValidator(commit.Body, config.Level, config.Value.(int))
	return
}

func bodyMaxLineLengthValidator(body string, level validation.ValidationState, value int) (rule string, state validation.ValidationState) {
	rule = fmt.Sprintf("Body can't have lines longer than %d characters", value)

	if len(body) == 0 {
		return
	}

	splitted := strings.SplitSeq(body, "\n")
	for line := range splitted {
		if len(line) > value {
			state = level
			break
		}
	}

	return
}
