package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	HEADER_MIN_LENGTH_SHORT_HEADER = "a"
	HEADER_MIN_LENGTH_LONG_HEADER  = "ab"
)

var (
	HEADER_MIN_LENGTH_EMPTY = parser.ParseConventionalCommit("").Header
	HEADER_MIN_LENGTH_SHORT = parser.ParseConventionalCommit(HEADER_MIN_LENGTH_SHORT_HEADER).Header
	HEADER_MIN_LENGTH_LONG  = parser.ParseConventionalCommit(HEADER_MIN_LENGTH_LONG_HEADER).Header
)

func TestHeaderMinLengthEmptyShouldSucceed(t *testing.T) {
	_, level := headerMinLengthValidator(
		HEADER_MIN_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderMinLengthShortShouldFail(t *testing.T) {
	_, level := headerMinLengthValidator(
		HEADER_MIN_LENGTH_SHORT,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderMinLengthLongShouldSucceed(t *testing.T) {
	_, level := headerMinLengthValidator(
		HEADER_MIN_LENGTH_LONG,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
