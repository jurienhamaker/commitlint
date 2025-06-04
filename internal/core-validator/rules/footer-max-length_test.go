package rules

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	FOOTER_MAX_LENGTH_SHORT_FOOTER = "Closes #1"
	FOOTER_MAX_LENGTH_LONG_FOOTER  = "Closes #1 and #2"
	FOOTER_MAX_LENGTH_MULTI_FOOTER = "Closes #1\nCloses #2"
)

var (
	FOOTER_MAX_LENGTH_EMPTY = parser.ParseConventionalCommit("test: subject").Footer
	FOOTER_MAX_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", FOOTER_MAX_LENGTH_SHORT_FOOTER)).Footer
	FOOTER_MAX_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", FOOTER_MAX_LENGTH_LONG_FOOTER)).Footer
	FOOTER_MAX_LENGTH_MULTI = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", FOOTER_MAX_LENGTH_MULTI_FOOTER)).Footer
)

func TestFooterMaxLengthEmptyShouldSucceed(t *testing.T) {
	_, level := footerMaxLengthValidator(
		FOOTER_MAX_LENGTH_EMPTY,
		validation.ValidationStateError,
		len(strings.Join(FOOTER_MAX_LENGTH_SHORT, "\n")),
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMaxLengthShortShouldSucceed(t *testing.T) {
	_, level := footerMaxLengthValidator(
		FOOTER_MAX_LENGTH_SHORT,
		validation.ValidationStateError,
		len(strings.Join(FOOTER_MAX_LENGTH_SHORT, "\n")),
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMaxLengthLongShouldFail(t *testing.T) {
	_, level := footerMaxLengthValidator(
		FOOTER_MAX_LENGTH_LONG,
		validation.ValidationStateError,
		len(strings.Join(FOOTER_MAX_LENGTH_SHORT, "\n")),
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMaxLengthMultiShouldFail(t *testing.T) {
	_, level := footerMaxLengthValidator(
		FOOTER_MAX_LENGTH_MULTI,
		validation.ValidationStateError,
		len(strings.Join(FOOTER_MAX_LENGTH_SHORT, "\n")),
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
