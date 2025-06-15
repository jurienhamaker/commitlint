package rules

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	TYPE_MIN_LENGTH_SHORT_TYPE = "a"
	TYPE_MIN_LENGTH_LONG_TYPE  = "ab"
)

var (
	TYPE_MIN_LENGTH_EMPTY = parser.ParseConventionalCommit("subject").Type
	TYPE_MIN_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("%s: subject", TYPE_MIN_LENGTH_SHORT_TYPE)).Type
	TYPE_MIN_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("%s: subject", TYPE_MIN_LENGTH_LONG_TYPE)).Type
)

func TestTypeMinLengthEmptyShouldSucceed(t *testing.T) {
	_, level := typeMinLengthValidator(
		TYPE_MIN_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeMinLengthShortShouldFail(t *testing.T) {
	_, level := typeMinLengthValidator(
		TYPE_MIN_LENGTH_SHORT,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeMinLengthLongShouldSucceed(t *testing.T) {
	_, level := typeMinLengthValidator(
		TYPE_MIN_LENGTH_LONG,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
