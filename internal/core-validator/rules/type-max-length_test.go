package rules

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	TYPE_MAX_LENGTH_SHORT_TYPE = "a"
	TYPE_MAX_LENGTH_LONG_TYPE  = "ab"
)

var (
	TYPE_MAX_LENGTH_EMPTY = parser.ParseConventionalCommit("type").Type
	TYPE_MAX_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("%s: subject", TYPE_MAX_LENGTH_SHORT_TYPE)).Type
	TYPE_MAX_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("%s: subject", TYPE_MAX_LENGTH_LONG_TYPE)).Type
)

func TestTypeMaxLengthEmptyShouldSucceed(t *testing.T) {
	_, level := typeMaxLengthValidator(
		TYPE_MAX_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeMaxLengthShortShouldSucceed(t *testing.T) {
	_, level := typeMaxLengthValidator(
		TYPE_MAX_LENGTH_SHORT,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeMaxLengthLongShouldFail(t *testing.T) {
	_, level := typeMaxLengthValidator(
		TYPE_MAX_LENGTH_LONG,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
