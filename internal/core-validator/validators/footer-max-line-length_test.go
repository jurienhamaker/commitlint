package validators

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER = "Closes #1"
	FOOTER_MAX_LINE_LENGTH_LONG_FOOTER  = "Closes #1 and #2"
)

var (
	FOOTER_MAX_LINE_LENGTH_EMPTY       = parser.ParseConventionalCommit("test: subject").Footer
	FOOTER_MAX_LINE_LENGTH_SHORT       = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER)).Footer
	FOOTER_MAX_LINE_LENGTH_LONG        = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", FOOTER_MAX_LINE_LENGTH_LONG_FOOTER)).Footer
	FOOTER_MAX_LINE_LENGTH_SHORT_MULTI = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s\n%s", FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER, FOOTER_MAX_LENGTH_SHORT_FOOTER)).Footer
	FOOTER_MAX_LINE_LENGTH_LONG_MULTI  = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s\n%s", FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER, FOOTER_MAX_LENGTH_LONG_FOOTER)).Footer
)

func TestFooterMaxLineLengthEmptyShouldSucceed(t *testing.T) {
	_, level := footerMaxLineLengthValidator(
		FOOTER_MAX_LINE_LENGTH_EMPTY,
		validation.ValidationStateError,
		len(FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER),
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMaxLineLengthShortShouldSucceed(t *testing.T) {
	_, level := footerMaxLineLengthValidator(
		FOOTER_MAX_LINE_LENGTH_SHORT,
		validation.ValidationStateError,
		len(FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER),
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMaxLineLengthLongShouldFail(t *testing.T) {
	_, level := footerMaxLineLengthValidator(
		FOOTER_MAX_LINE_LENGTH_LONG,
		validation.ValidationStateError,
		len(FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER),
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMaxLineLengthShortMultiShouldSucceed(t *testing.T) {
	_, level := footerMaxLineLengthValidator(
		FOOTER_MAX_LINE_LENGTH_SHORT_MULTI,
		validation.ValidationStateError,
		len(FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER),
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterMaxLineLengthLongMultiShouldFail(t *testing.T) {
	_, level := footerMaxLineLengthValidator(
		FOOTER_MAX_LINE_LENGTH_LONG_MULTI,
		validation.ValidationStateError,
		len(FOOTER_MAX_LINE_LENGTH_SHORT_FOOTER),
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
