package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	HEADER_MAX_LENGTH_SHORT_HEADER = "a"
	HEADER_MAX_LENGTH_LONG_HEADER  = "ab"
)

var (
	HEADER_MAX_LENGTH_EMPTY = parser.ParseConventionalCommit("").Header
	HEADER_MAX_LENGTH_SHORT = parser.ParseConventionalCommit(HEADER_MAX_LENGTH_SHORT_HEADER).Header
	HEADER_MAX_LENGTH_LONG  = parser.ParseConventionalCommit(HEADER_MAX_LENGTH_LONG_HEADER).Header
)

func TestHeaderMaxLengthEmptyShouldSucceed(t *testing.T) {
	_, level := headerMaxLengthValidator(
		HEADER_MAX_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderMaxLengthShortShouldSucceed(t *testing.T) {
	_, level := headerMaxLengthValidator(
		HEADER_MAX_LENGTH_SHORT,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderMaxLengthLongShouldFail(t *testing.T) {
	_, level := headerMaxLengthValidator(
		HEADER_MAX_LENGTH_LONG,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
