package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	FOOTER_EMPTY_EMPTY = parser.ParseConventionalCommit("feat:\nBody").Footer
	FOOTER_EMPTY_WITH  = parser.ParseConventionalCommit("test: subject\nbody\nCloses #1.").Footer
)

func TestFooterEmptyEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := footerEmptyValidator(
		FOOTER_EMPTY_EMPTY,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterEmptyEmptyNeverShouldFail(t *testing.T) {
	_, level := footerEmptyValidator(
		FOOTER_EMPTY_EMPTY,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterEmptyFilledAlwaysShouldFail(t *testing.T) {
	_, level := footerEmptyValidator(
		FOOTER_EMPTY_WITH,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterEmptyFilledNeverShouldSucceed(t *testing.T) {
	_, level := footerEmptyValidator(
		FOOTER_EMPTY_WITH,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
