package rules

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	FOOTER_MIN_LENGTH_SHORT_FOOTER = "Closes #1"
	FOOTER_MIN_LENGTH_LONG_FOOTER  = "Closes #1 and #2"
)

var (
	FOOTER_MIN_LENGTH_EMPTY = parser.ParseConventionalCommit("test: subject").Footer
	FOOTER_MIN_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", FOOTER_MIN_LENGTH_SHORT_FOOTER)).Footer
	FOOTER_MIN_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", FOOTER_MIN_LENGTH_LONG_FOOTER)).Footer
)

func TestFooterMinLengthEmptyShouldSucceed(t *testing.T) {
	_, level := footerMinLengthValidator(
		FOOTER_MIN_LENGTH_EMPTY,
		validation.ValidationStateError,
		len(FOOTER_MIN_LENGTH_LONG_FOOTER),
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMinLengthShortShouldFail(t *testing.T) {
	_, level := footerMinLengthValidator(
		FOOTER_MIN_LENGTH_SHORT,
		validation.ValidationStateError,
		len(FOOTER_MIN_LENGTH_LONG_FOOTER),
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMinLengthLongShouldSucceed(t *testing.T) {
	_, level := footerMinLengthValidator(
		FOOTER_MIN_LENGTH_LONG,
		validation.ValidationStateError,
		len(FOOTER_MIN_LENGTH_LONG_FOOTER),
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
