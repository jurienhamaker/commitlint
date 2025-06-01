package validators

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	BODY_MIN_LENGTH_SHORT_BODY = "a"
	BODY_MIN_LENGTH_LONG_BODY  = "ab"
)

var (
	BODY_MIN_LENGTH_EMPTY = parser.ParseConventionalCommit("test: subject").Body
	BODY_MIN_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", BODY_MIN_LENGTH_SHORT_BODY)).Body
	BODY_MIN_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", BODY_MIN_LENGTH_LONG_BODY)).Body
)

func TestBodyMinLengthEmptyShouldSucceed(t *testing.T) {
	_, level := bodyMinLengthValidator(
		BODY_MIN_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyMinLengthShortShouldFail(t *testing.T) {
	_, level := bodyMinLengthValidator(
		BODY_MIN_LENGTH_SHORT,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyMinLengthLongShouldSucceed(t *testing.T) {
	_, level := bodyMinLengthValidator(
		BODY_MIN_LENGTH_LONG,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
