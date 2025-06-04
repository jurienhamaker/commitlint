package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	BODY_LEADING_BLANK_SIMPLE  = parser.ParseConventionalCommit("test: subject").Body
	BODY_LEADING_BLANK_WITHOUT = parser.ParseConventionalCommit("test: subject\nbody.").Body
	BODY_LEADING_BLANK_WITH    = parser.ParseConventionalCommit("test: subject\n\nbody.").Body
)

func TestBodyLeadingBlankSimpleAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyLeadingBlankValidator(
		BODY_LEADING_BLANK_SIMPLE,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyLeadingBlankSimpleNeverShouldSucceed(t *testing.T) {
	_, level := bodyLeadingBlankValidator(
		BODY_LEADING_BLANK_SIMPLE,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyLeadingBlankWithoutAlwaysShouldFail(t *testing.T) {
	_, level := bodyLeadingBlankValidator(
		BODY_LEADING_BLANK_WITHOUT,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyLeadingBlankWithoutNeverShouldSucceed(t *testing.T) {
	_, level := bodyLeadingBlankValidator(
		BODY_LEADING_BLANK_WITHOUT,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyLeadingBlankWithAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyLeadingBlankValidator(
		BODY_LEADING_BLANK_WITH,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyLeadingBlankWithNeverShouldFail(t *testing.T) {
	_, level := bodyLeadingBlankValidator(
		BODY_LEADING_BLANK_WITH,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
